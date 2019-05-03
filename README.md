# Team Mavericks
Team project for CMPE-281, Cloud Computing <br>
Submitted To: [Prof Paul Nguyen](https://github.com/paulnguyen)

Team members:
1. [Arihant Sai](https://github.com/Arihant1467)
2. [Pratik Bhandarkar](https://github.com/pratikb25)
3. [Sayali Patil](https://github.com/SayaliPatil)
4. [Sharwari Phadnis](https://github.com/sharwari09)
5. [Thol Chidambaram](https://github.com/thol)

# Project Architecture
<img src="./images/cmpe281_arch.png"/>

# Project Journal
[Project Journal](ProjectJournal.md)

# Kanban Board
[Kanban Task Board](https://github.com/nguyensjsu/sp19-281-mavericks/projects/1)

# Team Meetings
[Team Meetings Log](ProjectJournal.md#Minutes-of-Meeting)

# Lambda vs Kubernetes

Serverless architectures — which in many ways is simply a repackaging and re-imagining of microservice architectures — is competing with Kubernetes because it allows for the scaling of applications and deployments without the complexity and configuration headaches of Kubernetes, or even containers. But don’t confuse the two as being equal.

A serverless application, if there are no requests for any of its functions, can drive costs to zero. they essentially cease to exist unless they are explicitly accessed. This can lead to dramatically lower costs, and also much faster scaling. The more a serverless application is accessed, the larger it scales.

The idea that serverless architectures will replace containerized applications does not seem to be a rational proposal. Not everything can be reduced to an ephemeral function. Some applications will always require the ability to persist data and state while an application is running, and this is not something that Serverless architectures are particularly designed for. But interest in Serverless is nevertheless growing rapidly.

[More information](https://thenewstack.io/why-serverless-vs-kubernetes-isnt-a-real-debate/)
