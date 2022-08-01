const scanner = require('sonarqube-scanner');
scanner(
  {
  serverUrl: "http://107.20.130.141:9000/",
  login:"admin",
  password:"sonar",
  options: {
    "sonar.sources": "./src"
  },
},
() => process.exit()
);
