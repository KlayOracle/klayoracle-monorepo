#!/usr/bin/env node
'use strict'

const options = require('minimist')(process.argv.slice(2))
const stdin = require('get-stdin')()
const {split,width} = require('./')

stdin.then(s => console.log(options.width
	? width(s)
	: split(s, options.columns || 80).join('\n')
), console.error)
