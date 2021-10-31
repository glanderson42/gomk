![Logo](assets/logo.png)

### About
---
Gomk is a json based build system for go applications written in go. You just have to define your targets and it is build the magic!

### Quick example of a ```gomk.json```
---
```
{
  "project": "sample-project",
  "installModules": false,
  "targets": [
    {
      "bin": "sample-target",
      "outputDir": "bin",
      "sourceDir": "sample-project",
      "flags": [
        "-race"
      ],
      "vendor": false,
      "release": false
    }
  ]
}
```

#### Generate Makefile from ```gomk.json```
```gomk --generate-makefile```

#### Generate example ```gomk.json```
```gomk --generate-sample```

### Installation
---
```
git clone https://github.com/glanderson42/gomk
cd gomk
# todo
```