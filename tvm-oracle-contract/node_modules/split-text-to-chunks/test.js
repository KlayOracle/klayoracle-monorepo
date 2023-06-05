'use strict'

const test = require('tape')
const {split,width} = require('./')

test(function( t ){
	t.deepEqual(split('A lazy dog', 1), a`
		A
		l
		a
		z
		y
		d
		o
		g
	`)
	t.deepEqual(split('A lazy dog', 2), a`
		A_
		la
		zy
		do
		g
	`)
	t.deepEqual(split('A lazy dog', 3), a`
		A l
		azy
		dog
	`)
	t.deepEqual(split('A lazy dog', 4), a`
		A_
		lazy
		dog
	`)
	t.deepEqual(split('A lazy dog', 5), a`
		A_
		lazy_
		dog
	`)
	t.deepEqual(split('A lazy dog', 6), a`
		A lazy
		dog
	`)
	t.deepEqual(split('A lazy dog', 7), a`
		A lazy_
		dog
	`)
	t.deepEqual(split('A lazy dog', 8), a`
		A lazy_
		dog
	`)
	t.deepEqual(split('A lazy dog', 9), a`
		A lazy_
		dog
	`)
	t.deepEqual(split('A lazy dog', 10), a`
		A lazy dog
	`)

	t.deepEqual(split('A lazy  dog', 5), a`
		A_
		lazy_
		dog
	`)
	t.deepEqual(split('A lazy   dog', 5), a`
		A_
		lazy_
		_dog
	`)
	t.deepEqual(split('A lazy    dog', 5), a`
		A_
		lazy_
		__dog
	`)
	t.deepEqual(split('A lazy     dog', 5), a`
		A_
		lazy_
		___
		dog
	`)
	t.deepEqual(split('     ', 2), a`
		__
		__
	`)

	t.deepEqual(split('A lazy 🐶 made a pile of 💩', 8), a`
		A lazy 🐶
		made a_
		pile of_
		💩
	`)

	t.deepEqual(split('A lazy\n\n\ndog', 15), a`
		A lazy


		dog
	`)

	t.deepEqual(split('Initial long word', 5), a`
		Initi
		al_
		long_
		word
	`)

	t.deepEqual(split('Word wrap', 10), a`
		Word wrap
	`, 'No wrapping')

	t.deepEqual(split('Word wrap', 4), a`
		Word
		wrap
	`, 'Soft wrap, full width words')

	t.deepEqual(split('Word wrap', 5), a`
		Word_
		wrap
	`, 'Soft wrap, width -1 words')

	t.deepEqual(split('Word wrap', 6), a`
		Word_
		wrap
	`, 'Soft wrap, width -2 words')

	t.deepEqual(split('Hard wrap', 3), a`
		Har
		d w
		rap
	`, 'Hard wrap')

	t.deepEqual(split('Words wrap', 4), a`
		Word
		s_
		wrap
	`, 'Prefer soft over hard wrap')

	t.deepEqual(split('Word  wrap', 4), a`
		Word
		_
		wrap
	`)

	t.deepEqual(split('Is it word-wrap?', 10), a`
		Is it_
		word-wrap?
	`, 'Symbols count')

	t.equal(width('A lazy 🐶\nmade a\npile'), 8)
	t.equal(width('A lazy 🐶\nmade a\npile', 6), 6)
	t.equal(width('A lazy 🐶\nmade a pile of 💩'), 16)
	t.equal(width('Žž'), 2)

	t.end()
})

function a(...lines){
	return lines.join('').trim().replace(/\t/g, '').replace(/_/g, ' ').split('\n')
}
