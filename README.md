<p align="center">
   <br/>
   <h3 align="center">The Open API </h3>
   <br/>
   <p align="center">
   netpalm makes it easy to push and pull state from your apps to your network by providing multiple southbound drivers, abstraction methods and modern northbound interfaces such as open API3 and REST webhooks.
   </p> 
   <p align="center" style="align: center;">
      <img src="https://github.com/tbotnz/netpalm/workflows/tests/badge.svg" alt="Tests"/>


# webex-meeting-types-samples

[![published](https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg)](https://testing-developer.cisco.com/codeexchange/github/repo/lingjshi/DNAC-Webex-Teams-Bot-App) [![Run in VSCode](https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-runable-icon.svg)](https://testing-developer.cisco.com/codespace/?id=devenv-vscode-base-webex-demo&community=true&type=vscode&GITHUB_SOURCE_REPO=https://github.com/lingjshi/webex-meeting-types-samples)

### Overview

Demonstrates 'webpack' bundling of the Webex JavaScript SDK and [Momentum-UI](https://github.com/momentum-design/momentum-ui/) style assets for use in browser voice/video meeting application integrations.  The resulting page/bundle is served via a simple light web server as a single-page app.

Includes examples of accessing/joining various meeting types, including:

* 1:1 Webex cloud calling
* Space multi-user cloud calling
* Scheduling and joining Webex scheduling meetings
* PMR meetings
* Webex Calling/PSTN dialing

Bundling framework demonstrated:

* [Webpack](https://webpack.js.org/)

>This project was built/tested using:

>* [Visual Studio Code](https://code.visualstudio.com/)
>* Ubuntu 22.04
>* Node 16.14.2

[https://developer.webex.com/docs/sdks/browser](https://developer.webex.com/docs/sdks/browser)

### Getting started

#### 1.Open VSCode DevEnv

The following repo has been cloned automatically:

```
https://github.com/CiscoDevNet/webex-meeting-types-samples.git
```

#### 2.In VS Code terminal, install dependencies

```bash
npm install
```

#### 3.Run the build command in VS Code terminal

```bash
npm run build
```

#### 4.Run the command to launch the webserver

```bash
npm run launch
```

#### 5.Open a new terminal and run the command to get the url

```bash
echo $DEVENV_APP_9082_URL
```
   
#### 6.Open the target page in your browser using the url

You can test the sample by logging into [developer.webex.com](https://developer.webex.com) and grabbing a Personal Access Token from the [Getting Started](https://developer.webex.com/docs/api/getting-started) page, then dialing another Webex Teams user via their Webex Id/email

>**Note:** Don't connect and dial based on the same user - that won't work!

### Hints

There is a workaround in webpack/webpack.config.js for an [issue](https://github.com/webpack-contrib/css-loader/issues/447) Webpack has with the `fs` module that's a dependency of `webex`, but not actually needed in browser usage:

```javascript
...
node: {
    fs: 'empty'
}
...
```

See `package.json` for the `browserlists` array of target browsers/versions

## Automated Use Case
Remote observance is especially useful for offices, for example, and remote observance is in demand due to the fact that we are transitioning towards a hybrid work experience. Moreover, in our use case, there is a bee hive located at the office. For safety reasons, bee hive inspections have to be done with two or more people. In the unlikely event of an accident, the second person can provide assistance or call the emergency services. However, since fewer people are attending the office, it is sometimes inconvenient to have two people doing the inspection. Therefore, we have created a process and integation that allows us to conduct remote observance, where one team member can do the bee hive inspection alone, while team members observe the site remotely and can contact assistance if needed.

In addition to remote observance, the team would also like to monitor the honey bee colony. For this PoV, we have installed three Meraki MV cameras and a BEEP base, which encompasses multiple sensors to monitor a bee hive colony. We have created a bot that would interact during the remote observance process, but we have added features that allows the users to generate snapshots from the Meraki cameras or query the latest metric, e.g. temperature, humidity, weight, etc.

Related Code Repo B:  [LInk B >](/codeexchange/github/repo/hhxiao/gve_devnet_meraki_alert_webex_bot_notification/)

Related Code Repo C:  [LInk C >](/codeexchange/github/repo/hhxiao/gve_devnet_meraki_alert_webex_bot_notification/)

Related Code Repo D:  [LInk D >](/codeexchange/github/repo/hhxiao/gve_devnet_meraki_alert_webex_bot_notification/)


## Learning Labs
[Introduction to Meraki Open API Specification documentation](https://developer.cisco.com/learning/labs/collab-webex-apps/)

## White Papers
[Cloud-Managed Smart Cameras: Meraki MV](https://developer.cisco.com/learning/labs/collab-webex-apps/)

## Sandbox
[Meraki Always On](https://devnetsandbox.cisco.com/RM/Topology)

