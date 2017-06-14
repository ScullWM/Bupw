Bupw
=========
Basic User Password Warning

### About
Bupw (no idea about how to pronounce it) is a simple "micro-service" application that check if a password is a minimal secure enougth.
Perfect tools to check if a term is present in the [Probable-Wordlists](https://github.com/berzerk0/Probable-Wordlists).


### Usage
First step is to configure the `config.yml` file
`./bupw`

Then try it using: `curl http://localhost:8000/\?password\=testmdfjkngops`
Reponse is json formated and return a true if term was found in a dictionnary, false if not found.


### Configuration
Here's the config.yml default values.
```
httpserver:
    port: 8000
    query: password
files:
  - top95.txt
  - wpa.txt
  - dico.txt

```

You may find various dictonnary file from [weakpass.com](https://weakpass.com/)

### Contribute
To contribute just open a Pull Request with your new code. Feel free to contribute.