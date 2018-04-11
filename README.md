# goYoda [![Donate](https://img.shields.io/badge/Donate-PayPal-blue.svg?style=for-the-badge)](https://www.paypal.me/HaoZeke/) [![GoDoc](https://godoc.org/github.com/HaoZeke/goYoda?status.svg)](https://godoc.org/github.com/HaoZeke/goYoda) [![Build Status](https://semaphoreci.com/api/v1/haozeke/docuyoda_starter/branches/master/badge.svg)](https://semaphoreci.com/haozeke/docuyoda_starter)

> Copyright (C) 2017  Rohit Goswami <rohit1995@mail.ru>

![](readme/turtle.png "Pandoc for turtles")

**Check the output [here](https://github.com/HaoZeke/docuYoda_Starter/blob/pdf/spooky-action.pdf).**

This is a consolidation of both the [docuYoda](https://github.com/HaoZeke/docuYoda_Starter) and [zenYoda](http://zenyoda.surge.sh/) projects.

As such this is to be distributed as a single binary written in go, which enables high level rendering of markdown or other documents powered by pandoc.

Both pandoc style citations and traditional TeX citations are enabled (according to the flags), however keep in mind that TeX citations will not display in non-TeX files.

Read about the project at it's source [here](https://www.github.com/HaoZeke/goYoda).

For creating presentations, refer to the sibling project, [zenYoda](http://zenyoda.surge.sh/) and it's [starter template](http://zenyodasap.surge.sh/).

## Goals

* Single binary
* Supports templates
* Easily extensible
* HTML based live pandoc preview
* Integration with surge or netlify for presentations
* Focus on presentations, reports and flashcards
* Interactive and file based configuration
* Continuous compilation
* Cross Platform
* Full test suite

## Acknowledgments

I am indebted to the following projects:

* [pandoc-viewer](https://github.com/EntilZha/pandoc-viewer) by EntilZha
* [go project layout](https://github.com/golang-standards/project-layout)
* [goreleaser](https://goreleaser.com/)
* [urfave/cli](https://github.com/urfave/cli)
* [packr](https://github.com/gobuffalo/packr)

For a full list of important projects look at the [Gopkg.toml](./Gopkg.toml) file.

## License
Refer to the project license.

The project like much of pandoc itself is under the [GNU GPLv3](https://choosealicense.com/licenses/gpl-3.0/), however, please refer to the exceptions listed [here](https://github.com/jgm/pandoc/blob/master/COPYRIGHT).