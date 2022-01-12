# SSM-CLI #

This is a simple golang CLI to enable users to easily start AWS Systems Manager 
Sessions from the terminal, without having to remember the command themselves 
or list instance IDs.

## Usage

With some enviroment variables for AWS set, just run `ssmcli`. The script will list existing sessions first so you can reconnect if one died, then it will list all instances.

Session Manager needs to be working for this script to work, it's just a wrapper! So if the agent isn't installed or can't connect to AWS, it won't work.

## Detailed Usage

If you've no profile or variables or token, you'll get this:

    $ ssmcli
    ERRO[0005] failed to list sessions                       error="NoCredentialProviders: no valid providers in chain. Deprecated.\n\tFor verbose messaging see aws.Config.CredentialsChainVerboseErrors"
    ERRO[0005] no credentials provider found, do you need to set AWS_* variables?
    ERRO[0005] no instances found

With variables set, you'll get a list of instances that Session Manager knows about:

    DEBU[0000] instance found                                id=i-078               name=shearn89.ec2.jenkins
    DEBU[0000] instance found                                id=i-040               name=shearn89.ec2.centos.7
    DEBU[0000] map[shearn89.ec2.centos.7:i-040               shearn89.ec2.jenkins:i-078              ]
    Use the arrow keys to navigate: ↓ ↑ → ←
    ? Select instance:
      ▸ shearn89.ec2.jenkins
        shearn89.ec2.centos.7

When you select one, you get the option to start a shell, or forward ports:

    DEBU[0000] instance found                                id=i-078               name=shearn89.ec2.jenkins
    DEBU[0000] instance found                                id=i-040               name=shearn89.ec2.centos.7
    DEBU[0000] map[shearn89.ec2.centos.7:i-040               shearn89.ec2.jenkins:i-078              ]
    ✔ shearn89.ec2.jenkins
    DEBU[0135] selected instance                             id=i-078               name=shearn89.ec2.jenkins
    Use the arrow keys to navigate: ↓ ↑ → ←
    ? Select action:
      ▸ start SSH
        forward ports

If you then select SSH, you get a session:

    DEBU[0000] instance found                                id=i-078               name=shearn89.ec2.jenkins
    DEBU[0000] instance found                                id=i-040               name=shearn89.ec2.centos.7
    DEBU[0000] map[shearn89.ec2.centos.7:i-040               shearn89.ec2.jenkins:i-078              ]
    ✔ shearn89.ec2.jenkins
    DEBU[0135] selected instance                             id=i-078               name=shearn89.ec2.jenkins
    ✔ start SSH
    DEBU[0170] selected Action                               action="start SSH"
    INFO[0170] starting SSH shell
    
    Starting session with SessionId: shearn89-0a0              
    sh-4.2$

If you want to forward ports, then you get the following:

    DEBU[0000] instance found                                id=i-078               name=shearn89.ec2.jenkins
    DEBU[0000] instance found                                id=i-040               name=shearn89.ec2.centos.7
    DEBU[0000] map[shearn89.ec2.centos.7:i-040               shearn89.ec2.jenkins:i-078              ]
    ✔ shearn89.ec2.jenkins
    DEBU[0001] selected instance                             id=i-078               name=shearn89.ec2.jenkins
    ✔ forward ports
    DEBU[0002] selected Action                               action="forward ports"
    Use the arrow keys to navigate: ↓ ↑ → ←
    ? Select port to forward:
      ▸ 22
        80
        443
        8080

(TODO?) And then selecting the port runs the script:

    DEBU[0000] instance found                                id=i-078               name=shearn89.ec2.jenkins
    DEBU[0000] instance found                                id=i-040               name=shearn89.ec2.centos.7
    DEBU[0000] map[shearn89.ec2.centos.7:i-040               shearn89.ec2.jenkins:i-078              ]
    ✔ shearn89.ec2.jenkins
    DEBU[0001] selected instance                             id=i-078               name=shearn89.ec2.jenkins
    ✔ forward ports
    DEBU[0002] selected Action                               action="forward ports"
    ✔ 8080
    DEBU[0022] selected port                                 port=8080
    INFO[0022] run port forwarding                           port=8080

## Building

Needs a version of Go that supports modules, and then:

    go install 

Put it somewhere on your PATH or add GOBIN to your PATH.
