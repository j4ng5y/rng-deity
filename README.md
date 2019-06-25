[![CircleCI](https://circleci.com/gh/j4ng5y/rng-deity/tree/releases%2F2019.06.24.svg?style=svg)](https://circleci.com/gh/j4ng5y/rng-deity/tree/releases%2F2019.06.24)
[![License](https://img.shields.io/github/license/j4ng5y/rng-deity.svg)](https://github.com/j4ng5y/rng-deity/tree/master/LICENSE.txt)
![RNG-DEITY](images/rng-deity.png)

# RNG-DEITY

In an attempt to be more inclusive in my Slack escapades, I wrote this to fill in the blank when I say "Thank {deity}" or something else of the sort.

This is a simple little app that outputs a single, random deity (in the format of "{NAME}, the {RELIGION} {TITLE} of {RANDOM GOD OVERSIGHT ITEM}") from a number of religions, including:

* Aleph
* Amish [X]
* Asatru [X]
* Bahai [X]
* Brahma Kumari [X]
* The Branch Dividians [X]
* Buddhist [X]
* Celtic [X]
* Chen Tao [X]
* Chinese [X]
* Christian [X]
* Christadelphians [X]
* Concernd Christians [X]
* Dami Xuanjiao [X]
* Devine Lightmission [X]
* Druidism [X]
* Druze [X]
* Dudeism [X]
* Ebionites [X]
* Eckankar [X]
* Egyptian [X] -- Major Deities Only, maybe I'll add the minor ones later
* Family of God [X]
* Gnosticism [X]
* Greek
* Hare Krishna
* Heathenism
* Heavens Gate
* Hikari No Wa
* Hinduism [X]
* Hookers for Jesus
* Humanism
* Islam
* Jainism
* Jedi
* Jehovah's Witnesses
* Judaism
* Mennonite
* Mesopotamian
* Mithraism
* Mysticism
* Native American
* New Age
* Norse
* Order of the Solar Temple
* Paganism
* Pastafaranism
* The Peoples Temple
* Raja Yoga
* Rastafarian
* Ravidassia
* Roman
* Salvation Army
* Santeria
* Satanic Temple
* Satanism
* Scientology
* Shamanism
* Shinto
* Sikhism
* Sumarian
* Taoism
* Thelema
* Traditional Africian
* Unifcation
* Unitaranism
* Unitarian Universalism
* Unitas Fratrum
* Voodoo
* Wicca
* Yezidism
* Zhu Shen Jiao
* Zoroastrianism

If the item has an "X" after it, I'm calling it done for now as I don't really have a TON of time to research all of this stuff.

Feel free to contribute. There is a file called "[template.go](https://github.com/j4ng5y/rng-deity/tree/master/deitylib/template.go)" that has a super basic template on how I wrote these things.

## Usage

```bash
A little app that outputs a random deity from a number of religions.

Usage:
  rng-deity [flags]

Flags:
  -h, --help   help for rng-deity (This message)
```
