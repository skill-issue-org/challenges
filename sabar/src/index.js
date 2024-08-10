/*
* Skill Issue Weekly Challenge
*
* Sabar
*
* Points: 4
* Author: dwisiswant0
*/

const cs = require('./consts'), rl = require('node:readline'),
  re = /^https?:\/\/(www\.)?([a-z0-9-]+)\.[a-z]{2,}\/([\w-]+)*$/;

const c = console, p = process,
  r = rl.createInterface({ input: p.stdin, output: p.stdout });

c.log(cs.banner)

r.question(cs.input + ': ', (async (i) => {
  try {
    const d = { redirect: 'manual' },
      f = await fetch(i + '/./../.../..../flag.txt', d), h = await f.text(),
      e = cs.error, n = (new Date).getTime(), o = cs.result + ': ' + cs.flag;

    re.test(i), (new Date).getTime() - n >= 5 ? c.log(o, h) : c.error(e);
  } catch(e) { c.error(e.message) }

  r.close()
}))
