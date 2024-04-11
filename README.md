![banner-symfony](./screenshots/homepage.png)

## Golang Gin Hackathon Starter (IN PROGRESS)

<!-- **Live Demo**: [https://symfony-hackathon-starter.herokuapp.com/](https://symfony-hackathon-starter.herokuapp.com/) -->

A boilerplate for **Golang Gin** web applications.

If you have attended any hackathons in the past, then you know how much time it takes to get a project started: decide on what to build, pick a programming language, pick a web framework, pick a CSS framework. A while later, you might have an initial project up on GitHub and only then can other team members start contributing. Or how about doing something as simple as Sign in with Facebook authentication? You can spend hours on it if you are not familiar with how OAuth 2.0 works.

The aim and usage of this project goes beyond hackathon purposes only, Golang Gin Hackathon starter will definitely save any developer hours or even days of development time and can serve as a learning guide for web developers in general.

[Golang](https://go.dev/) is an open-source programming language supported by Google. Software developers use Go in an array of operating systems and frameworks to develop web applications, cloud and networking services, and other types of software.

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

Gin Golang Hackathon Starter is a boilerplate application developed to keep you ahead in hackathons.

<h4 align="center">Modern Theme</h4>  
  
![](https://user-images.githubusercontent.com/19610753/47153649-1ad90800-d2d8-11e8-9f89-3164cfe11479.png)  
  
<h4 align="center">LOGIN</h4>  
  
![](https://user-images.githubusercontent.com/19610753/47153651-1b719e80-d2d8-11e8-90b1-f45cd9e08834.png)  
  
  
  <h4 align="center">REGISTER</h4>  
  
![](https://user-images.githubusercontent.com/19610753/47153652-1c0a3500-d2d8-11e8-85ff-ece0ec18d1e6.png)

Table of Contents

---

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Obtaining API Keys](#obtaining-api-keys)
- [Project Structure](#project-structure)
- [List of Packages](#list-of-packages)
- [Useful Tools and Resources](#useful-tools-and-resources)
- [Recommended Design Resources](#recommended-design-resources)
- [Recommended Symfony Libraries](#recommended-symfony-libraries)
- [Pro Tips](#pro-tips)
- [FAQ](#faq)
- [How It Works](#how-it-works-mini-guides)
- [Symfony Best Practises](#symfony-best-practises)
- [Deployment](#deployment)
- [Changelog](#changelog)
- [Contributing](#contributing)
- [License](#license)

Features

---

- **Local Authentication** using Email and Password
- **OAuth 1.0a Authentication** via Twitter
- **OAuth 2.0 Authentication** via Facebook, Google, GitHub, GitLab, LinkedIn, Auth0, Instagram
- Flash notifications
- MVC Project Structure
- Bootstrap 4
- Contact Form (powered by Gmail, Mailgun)
- **Account Management**
- Profile Details
- Change Password
- Forgot Password
- Reset Password
- Delete Account
- CSRF protection

Prerequisites

---

- [Mysql or Postgresql](https://www.mysql.com/ or http://www.postgresql.org/)
- [PHP 7.1.3+](http://php.net/)
- Command Line Tools
- <img src="http://deluge-torrent.org/images/apple-logo.gif" height="17">&nbsp;**Mac OS X:** [Xcode](https://itunes.apple.com/us/app/xcode/id497799835?mt=12) (or **OS X 10.9+**: `xcode-select --install`)
- <img src="http://dc942d419843af05523b-ff74ae13537a01be6cfec5927837dcfe.r14.cf1.rackcdn.com/wp-content/uploads/windows-8-50x50.jpg" height="17">&nbsp;**Windows:** [Visual Studio](https://www.visualstudio.com/products/visual-studio-community-vs)
- <img src="https://lh5.googleusercontent.com/-2YS1ceHWyys/AAAAAAAAAAI/AAAAAAAAAAc/0LCb_tsTvmU/s46-c-k/photo.jpg" height="17">&nbsp;**Ubuntu** / <img src="https://upload.wikimedia.org/wikipedia/commons/3/3f/Logo_Linux_Mint.png" height="17">&nbsp;**Linux Mint:** `sudo apt-get install build-essential`
- <img src="http://i1-news.softpedia-static.com/images/extra/LINUX/small/slw218news1.png" height="17">&nbsp;**Fedora**: `sudo dnf groupinstall "Development Tools"`
- <img src="https://en.opensuse.org/images/b/be/Logo-geeko_head.png" height="17">&nbsp;**OpenSUSE:** `sudo zypper install --type pattern devel_basis`

**Note:** If you are new to Symfony 4, I recommend to watch  
[Learn Symfony 4](https://symfonycasts.com/tracks/symfony), the official screencast for Symfony by Ryan Weaver and the SymfonyCasts crew. Alternatively,  
here is another great tutorial for building a project time tracker app for beginners/intermediate developers - [Let’s Build a Project Time Tracker with Symfony and VueJS](https://www.cloudways.com/blog/time-tracking-system-php-symfony-vue/).

Getting Started

---

#### Via Cloning The Repository:

The easiest way to get started is to clone the repository:

```bash
# Get the project
git clone https://github.com/yemiwebby/golang-gin-hackathon-starter.git
```

#### Change directory

```bash
cd golang-gin-hackathon-starter
```

#### Install dependencies

```bash
go mod tidy
```

#### Run the project

```bash
go run main.go
```

![Homepage](./screenshots/homepage.png)

This starter pack includes the following APIs. You will need to obtain appropriate credentials like Client ID, Client secret, API key, or Username & Password by going through each provider and generate new credentials.

- GitHub
- Twitter
- Auth0
- LinkedIn

Obtaining API Keys

---

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/2/2f/Google_2015_logo.svg/1000px-Google_2015_logo.svg.png" width="200">

- Visit [Google Cloud Console](https://cloud.google.com/console/project)
- Click on the **Create Project** button
- Enter _Project Name_, then click on **Create** button
- Then click on _APIs & auth_ in the sidebar and select _API_ tab
- Click on **Google+ API** under _Social APIs_, then click **Enable API**
- Next, under _APIs & auth_ in the sidebar click on _Credentials_ tab
- Click on **Create new Client ID** button
- Select _Web Application_ and click on **Configure Consent Screen**
- Fill out the required fields then click on **Save**
- In the _Create Client ID_ modal dialog:
- **Application Type**: Web Application
- **Authorized Javascript origins**: http://localhost:8000
- **Authorized redirect URI**: http://localhost:8000/secured/login_google
- Click on **Create Client ID** button
- Copy and paste _Client ID_ and _Client secret_ keys into `.env`

**Note:** When you ready to deploy to production don't forget to  
add your new url to _Authorized Javascript origins_ and _Authorized redirect URI_,  
e.g. `http://my-awesome-app.herokuapp.com` and  
`http://my-awesome-app.herokuapp.com/secured/login_google` respectively.  
The same goes for other providers.

<hr>

<img src="https://github.global.ssl.fastly.net/images/modules/logos_page/GitHub-Logo.png" width="200">
  
- Go to [Account Settings](https://github.com/settings/profile)  
- Select **Applications** from the sidebar  
- Then inside **Developer applications** click on **Register new application**  
- Enter *Application Name* and *Homepage URL*  
- For *Authorization Callback URL*: http://localhost:8000/secured/login_github  
- Click **Register application**  
- Now copy and paste *Client ID* and *Client Secret* keys into `.env` file  
  
<hr>  
  
<img src="https://g.twimg.com/ios_homescreen_icon.png" width="90">

- Sign in at [https://apps.twitter.com/](https://apps.twitter.com/)
- Click **Create a new application**
- Enter your application name, website and description
- For **Callback URL**: http://127.0.0.1:8000/secured/login_twitter
- Go to **Settings** tab
- Under _Application Type_ select **Read and Write** access
- Check the box **Allow this application to be used to Sign in with Twitter**
- Click **Update this Twitter's applications settings**
- Copy and paste _Consumer Key_ and _Consumer Secret_ keys into `.env` file

<hr>  
  
<img src="http://www.danpontefract.com/wp-content/uploads/2014/02/logo-linkedin.png" width="200">
  
- Sign in at [LinkedIn Developer Network](https://developer.linkedin.com/)  
- From the account name dropdown menu select **API Keys**  
 - *It may ask you to sign in once again*  
- Click **+ Add New Application** button  
- Fill out all the *required* fields  
 - **OAuth 2.0 Redirect URLs**: http://localhost:8000/secured/login_linkedin  
 - **JavaScript API Domains**: http://localhost:3000  
- For **Default Application Permissions** make sure at least the following is checked:  
 - `r_basicprofile`  
- Finish by clicking **Add Application** button  
- Copy and paste *API Key* and *Secret Key* keys into `.env` file  
 - *API Key* is your **clientID**  
 - *Secret Key* is your **clientSecret**  
  
<hr>  
  
<img src="https://raw.github.com/mailgun/media/master/Mailgun_Primary.png" width="200">

- Go to http://www.mailgun.com
- Sign up and add your _Domain Name_
- From the domain overview, copy and paste the default SMTP _Login_ and _Password_ into `.env` file

<hr>  
 
  
Project Structure  
-----------------  
  
| Name                                     | Description                                                  |  
| ----------------------------------       | ------------------------------------------------------------ |  
| **config**/config/fos_user.yaml                       | Configuration for FOS User Bundle              |  
| **config**/config/hwi_oauth.yaml                      | Configuration for HWIOAuthBundle                            |  
| **config**/config/security.yaml              | Configuration for Symfony security's system                               |  
| **config**/config/routes/routes.yaml                     | Configuration file for routing                |  
| **config**/config/services.yaml                  | This file is the entry point to configure your own services                                |  
| **controllers**/src/Controller/AccountController.php    | Controller for Account management                            |  
| **controllers**/src/Controller/ContactController.php    | Controller for Contact page                                  |  
| **controllers**/src/Controller/HomeController.php           | Controller for rendering all pages                                               |  
| **controllers**/user.js                  | Controller for user account management.                      |  
| **Entity**/src/Entity/User.php                      | User's Entity.                                             |  
| **Repository**/src/Repository/UserRepository.php                              | User Repository.                         |  
| **Provider**/src/Security/MyFOSUBProvider.php                              | Friends of Symfony User bundle provider.                         |  
| **Provider**/src/Auth0ResourceOwner.php                              | Provider for Auth0.                         |  
| **public**/                              | Static assets (fonts, css, js, img).                         |  
| **public**/**css**/style.css              | Main stylesheet for your app.                                |  
| **templates/bundles/FOSUserBundle**/             | This folder contains templates to override the default templates of FOSUserBundle for *login, password reset, signup, profile*.      |  
| **templates/default/index.html.twig**/                           | Templates for the homepage.                                  |  
| **templates/emails/contact.html.twig**/                           | Templates file for contact page form.                                  |  
| **templates/includes/footer.html.twig**/                           | Footer partial template.                                  |  
| **templates/includes/header.html.twig**/                           | Header partial template.                                  |  
| **templates/user/account.html.twig**/                           | Templates for the account page.                                  |  
| **templates/base.html.twig**/                           | Base template.                                  |  
| **templates/contact.html.twig**/                           | Templates for the contact page.                                  |  
| Procfile                              | Configuration file to set the web root on Heroku             |  
| .env.dist                             | Your API keys, tokens, passwords and database URI.           |  
| composer.json                            | File for loading all php packages.                           |  
| package.json                             | File for loading all necessary node modules.                 |  
  
  
List of Packages  
----------------  
  
| Package                         | Description                                                           |  
| ------------------------------- | --------------------------------------------------------------------- |  
| FOSUserBundle                       | Providers user management for Symfony project                             |  
| HWIOAuthBundle                       | OAuth client integration for Symfony. Supports both OAuth1.0a and OAuth2.                             |  
| phpunit                         | PHP testing library                                                   |  
| guzzlehttp                      | Simplified HTTP Request library                                       |  
  
  
Useful Tools and Resources  
--------------------------  
- [Symfony Blog](https://symfony.com/blog/) - Official Symfony blog  
- [Cloudways Blog](https://www.cloudways.com/blog/?s=symfony) - Symfony, PHP and other tutorials.  
- [SymfonyCasts](https://symfonycasts.com/) - Video tutorials for Symfony  
  
Recommended Design Resources  
----------------------------  
- [Code Guide](http://codeguide.co/) - Standards for developing flexible, durable, and sustainable HTML and CSS.  
- [Bootsnipp](http://bootsnipp.com/) - Code snippets for Bootstrap.  
- [UIBox](http://www.uibox.in) - Curated HTML, CSS, JS, UI components.  
- [Bootstrap Zero](https://www.bootstrapzero.com) - Free Bootstrap templates themes.  
- [Google Bootstrap](http://todc.github.io/todc-bootstrap/) - Google-styled theme for Bootstrap.  
- [Colors](http://clrs.cc) - A nicer color palette for the web.  
- [Creative Button Styles](http://tympanus.net/Development/CreativeButtons/) - awesome button styles.  
- [Creative Link Effects](http://tympanus.net/Development/CreativeLinkEffects/) - Beautiful link effects in CSS.  
- [Medium Scroll Effect](http://codepen.io/andreasstorm/pen/pyjEh) - Fade in/out header background image as you scroll.  
- [GeoPattern](https://github.com/btmills/geopattern) - SVG background pattern generator.  
- [Trianglify](https://github.com/qrohlf/trianglify) - SVG low-poly background pattern generator.  
 
  
FAQ  
--- 
  
How It Works (mini guides)  
--------------------------  
  
This section is intended for giving you a detailed explanation about how a particular functionality works. Maybe you are just curious about how it works, or maybe you are lost and confused while reading the code.
 
I hope it provides some guidance to you.  
  
<hr>  
  
### How do flash messages work in this project?  
Flash messages allow you to display a message at the end of the request and access it on next request and only next request. For instance, on a failed login attempt, you would display an alert with some error message, but as soon as you refresh that page or visit a different page and come back to the login page, that error message will be gone. It is only displayed once.  
  
<hr>  
  
### How do I create a new page?

Go through the following guide in order to successfully create a new page. [Check here](https://symfony.com/doc/current/page_creation.html).

Alternatively, you can also leverage the [Symfony maker bundle](https://symfony.com/doc/current/bundles/SymfonyMakerBundle/index.html) to generate a controller. Once you do this, it
will also automatically create a new templates file that can easily be rendered by the newly generated controller.

To begin, run the following command to have it installed via composer:

```bash
$ composer require symfony/maker-bundle --dev
```

If you are using Symfony hackathon starter project, you don't need to install the maker-bundle again. Next,
proceed to your terminal and run:

```bash
$ php bin/console make:controller SampleController
```

The preceding command will create two new file:

1. src/Controller/SampleController.php - the new controller
2. templates/sample/index.html.twig - `SampleController` templates file

To view the new page, ensure that your application is running (`php bin/console server:run`) and
navigate to `https://localhost:8000/sample`.

<hr>  
  
Symfony Cheatsheet  
-------------------  
  
* [Symfony Cheatsheet](https://www.cheatography.com/mika56/cheat-sheets/symfony/)

## Symfony Best Practises

- [Symfony Best Practises](https://symfony.com/doc/current/best_practices/index.html)

Deployment

---

Once you are ready to deploy your app, you will need to create an account with a cloud platform to host it. These are not the only choices, but they are my top  
picks. From my experience, **Heroku** is the easiest to get started with, deployments and custom domain support on free accounts.

### 1-Step Deployment with Heroku

<img src="http://blog.exadel.com/wp-content/uploads/2013/10/heroku-Logo-1.jpg" width="200">  
- Download and install [Heroku Toolbelt](https://toolbelt.heroku.com/)  
- In terminal, run `heroku login` and enter your Heroku credentials  
- From *your app* directory run `heroku create`  
- Create a Procfile in your app root. All this file needs to contain is `web: vendor/bin/heroku-php-nginx public` or `web: vendor/bin/heroku-php-apache2 public` if you prefer to use nginx.  
- Run `heroku addons:add heroku-postgresql:dev  ` to add a Postgres database to your heroku app from your terminal  
- Lastly, do `git push heroku master`.  Done!  
  
**Note:** To install Heroku add-ons your account must be verified.  
  
---  
  
- And you are done!  
  
<img src="https://cdn.iconscout.com/icon/free/png-256/digitalocean-1-283338.png" width="200" h>

This [tutorial](https://www.digitalocean.com/community/tutorials/how-to-deploy-a-symfony-4-application-to-production-with-lemp-on-ubuntu-18-04) will guide you on how to successfully deploy a Symfony 4 Application on Digital Ocean

## Contributing

Thank you for considering contributing to Symfony Hackathon Starter. The contribution guide can be found in the [Contribution File](CONTRIBUTING.md)

## Security Vulnerabilities

If you discover a security vulnerability within Symfony Hackathon Starter, please send an e-mail to Olususi Oluyemi at yemiwebby@gmail.com. All security vulnerabilities will be promptly addressed.

## Credits

- [Prosper Otemuyiwa](https://github.com/unicodeveloper/laravel-hackathon-starter) - Larathon
- [Sahat Yalkabov](https://github.com/sahat/hackathon-starter) - Awesome
- [Oluyemi Olususi](https://github.com/yemiwebby/symfony-hackathon-starter) - Symfony Hackathon Starter

## How can I thank you?

Why not star the github repo? I'd love the attention! Why not share the link for this repository on Twitter or HackerNews? Spread the word!

Don't forget to [follow me on twitter](https://twitter.com/yemiwebby)!

Thanks!  
Olususi Oluyemi.

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
