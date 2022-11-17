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

