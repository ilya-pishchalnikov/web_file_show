# Web File Show
This project enables you to synchronize a local text file and view it as a web page over HTTPS, allowing access from anywhere. The HTML page is formatted in the style of the VSCode Todo+ extension by Fabio Spampinatщ.
## Description
The project consists of two parts:
- **websrvfileshow** — a simple, lightweight web server supporting TLS and Basic Authentication. It offers methods to retrieve file content via POST requests and to display the content as an HTML page styled after the VSCode Todo+ extension by Fabio Spampinati.
- **filesender** - a lightweight background application for sending file content to a web server via POST, with customizable periodicity.
## Compile
1. Change directory to websrvfileshow or filesender to сompile project accordingly
2. To build for a different operating system, set the GOOS and GOARCH environment variables accordingly
3. Run "go build"
4. The executable will appear in the directory.
## Configuration
### websrvfileshow
Configuration is set through the config.json file located in the application directory and the websrvfileshow_credentials environment variable. 
#### config.json
config.json must be a valid json file containing single object with the following elements:
- **port** - the port which web server will listen
- **cert** - path to certificate *.pem file
- **certKey** - path to certificate key *.pem file
##### Example
```json
{
"port": ":443",
"cert": "cert.pem",
"certKey": "key.pem"
}
```
#### websrvfileshow_credentials environment variable
websrvfileshow_credentials environment variable contains credentials for web server authentification in format "user1=password1;user2=password2;...". Using ; or = characters in usernames or passwords is not allowed.
###filesender
Configuration is set through the config.json file located in the application directory
config.json must be a valid json file containing single object with the following elements:
- **fileName** - path to file which will be send to the web server
- **postMethod** - POST method address to send file content
- **periodicity** - file sending periodicity in seconds
- **userName** - user name for requesting web server
- **password** - password for requesting web server
##### Example
```json
{
    "fileName": "test.txt",
    "postMethod": "https://localhost/post_file",
    "periodicity": 1,
    "userName": "admin",
    "password": "Password123"
}