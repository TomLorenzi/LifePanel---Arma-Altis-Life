# LifePanel
Online panel for the base Altis Life framework

## EN

### How to install

#### Golang

What you need it GoLang and a website server nginx (or other)

If you want to install you can directly go there and follow the instructions : https://golang.org/doc/install

Here are my instructions with my version of the toolset 1.14.4 on Debian (should work on Ubuntu)

Go to the directory you want to download your files and type

```bash
curl https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz
```

Then type

```bash
sudo tar -C /usr/local -xzf go1.14.4.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
source /etc/profile
```

You've just installed GO

#### Pull for github

Then clone the repository where you need to place your files

#### Nginx config

This is only for nginx users but here is my configuration with my domain name that is Altasia.fr

```nginx
server {
        listen 80;
        server_name _;
        return 444;
}
server {
        listen 80;
        server_name www.altasia.fr altasia.fr;

        return 301 https://www.altasia.fr$request_uri;
}
server {
        listen 80;
        listen 443 ssl;
        server_name www.altasia.fr altasia.fr;

        ssl_certificate /etc/letsencrypt/live/www.altasia.fr/fullchain.pem;
        ssl_certificate_key /etc/letsencrypt/live/www.altasia.fr/privkey.pem;
        ssl_certificate /etc/letsencrypt/live/altasia.fr/fullchain.pem;
        ssl_certificate_key /etc/letsencrypt/live/altasia.fr/privkey.pem;

        location / {
                root /var/www/altasia;
        }
        
        location /api/ {
                proxy_pass http://127.0.0.1:8080/;
        }
}
```

#### Configure your API

Go to LifePanel/API/api.go and modify these lines :

```go
DBHOST string = "<ip adress of your DB server>"
DBPORT string = "3306" // Default port
DBUSER string = "<User>"
DBPWRD string = "<Password>"
DBNAME string = "altislife" //Should be correct
```

#### Configure your database

Run the file LifePanel.sql to add what's needed to your arma database

**WARNING**

If you changed the name of your database to something else than altislife you'll have to modify the API and the LifePanel.sql to be working !

### Run

To run your API go to the API folder and type : (The best way is to do it in a screen)

```bash
screen -S api
```

```bash
go build
./api
```

*To leave the screen type CTRL+A and D*

### Connection

You can identify yourself to the panel with this ID :
Login : Fondator
Password : password

**Please change it after the first connection !**

## FR

### Comment installer

#### Golang

Vous aurez besoin du toolset golang ainsi que d'un serveur web nginx (ou autre)

Si vous voulez installer golang vous pouvez directement suivre ces instructions : https://golang.org/doc/install

Voici mes instructions pour l'installation du toolset version 1.14.4 sur Debian (devrait fonctionner sur Ubuntu)

Rendez vous dans le dossier ou vous voulez télécharger le toolset et tapez :

```bash
curl https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz
```

Puis :

```bash
sudo tar -C /usr/local -xzf go1.14.4.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
source /etc/profile
```

Vous avez installé GoLang

#### Pull depuis Github

Clonez le repertoire la ou vous voulez qu'il soit (emplacement de vos fichiers web)

#### Nginx config

Cela concerne uniquement les utilisateurs de nginx, voici ma configuration avec comme nom de domain Altasia.fr

```nginx
server {
        listen 80;
        server_name _;
        return 444;
}
server {
        listen 80;
        server_name www.altasia.fr altasia.fr;

        return 301 https://www.altasia.fr$request_uri;
}
server {
        listen 80;
        listen 443 ssl;
        server_name www.altasia.fr altasia.fr;

        ssl_certificate /etc/letsencrypt/live/www.altasia.fr/fullchain.pem;
        ssl_certificate_key /etc/letsencrypt/live/www.altasia.fr/privkey.pem;
        ssl_certificate /etc/letsencrypt/live/altasia.fr/fullchain.pem;
        ssl_certificate_key /etc/letsencrypt/live/altasia.fr/privkey.pem;

        location / {
                root /var/www/altasia;
        }
        
        location /api/ {
                proxy_pass http://127.0.0.1:8080/;
        }
}
```

#### Configurer votre API

Ouvrez le fichier LifePanel/API/api.go et modifiez ces lignes

```go
DBHOST string = "<ip adress of your DB server>"
DBPORT string = "3306" // Default port
DBUSER string = "<User>"
DBPWRD string = "<Password>"
DBNAME string = "altislife" //Should be correct
```

#### Configurer votre BDD

Lancer le fichier LifePanel.sql (ou ses instructions directement dans votre logiciel de gestion de BDD) pour ajouter les tables necessaires

**ATTENTION**

Si vous avez changé le nom de votre BDD pour autre chose que altislife vous devrez modifier votre API et le fichier LifePanel.sql pour faire fonctionner le tout !

### Run

Pour démarrer votre API tapez : (le mieux est de le faire dans un screen)

```bash
screen -S api
```

```bash
go build
./api
```

*Pour quitter le screen tapez CTRL+A puis D*

### Connexion

Les identifiants lors de la première connexion sont :
Login : Fondator
Mot de passe : password

**Changez les dès la première connexion !**

## Credits

I worked from the base of a school project that i modified for myself and this project.

I used https://github.com/robinjulien/Leloux from https://github.com/robinjulien

Login Page and Home page are from a Template From ColorLib modified by me.

## Screens

<img src="screens\index.PNG" alt="index" style="zoom: 33%;" />

<img src="screens\player_list.PNG" alt="index" style="zoom: 50%;" />

<img src="screens\player_modify.PNG" alt="index" style="zoom: 50%;" />

<img src="screens\user_modify.PNG" alt="index" style="zoom: 50%;" />
