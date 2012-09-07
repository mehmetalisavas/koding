log = -> logger.info arguments...

# Core Nodejs libraries:
{spawn, exec}   = require 'child_process'
# crypto          = require 'crypto'
# sys             = require 'sys'
fs              = require 'fs'
Path            = require 'path'
{EventEmitter}  = require 'events'

slice         = Array::slice
splice        = Array::splice
noop          = Function()

# Error.stackTraceLimit = 100

if process.argv[5] is "true"
  __runCronJobs   = yes
  log "--cron is active, cronjobs will be running with your server."


process.on 'uncaughtException', (err)->
  exec './beep'
  console.log err, err?.stack


dbCallback= (err)->
  if err
    log err
    log "database connection couldn't be established - abort."
    process.exit()

if require("os").platform() is 'linux'
  require("fs").writeFile "/var/run/node/koding.pid",process.pid,(err)->
    if err?
      console.log "[WARN] Can't write pid to /var/run/node/kfmjs.pid. monit can't watch this process."

dbUrl = switch process.argv[3] or 'local'
  when "local"
    "mongodb://localhost:27017/koding?auto_reconnect"
  when "sinan"
    "mongodb://localhost:27017/kodingen?auto_reconnect"
  when "vpn"
    "mongodb://kodingen_user:Cvy3_exwb6JI@10.70.15.2:27017/kodingen?auto_reconnect"
  when "beta"
    "mongodb://beta_koding_user:lkalkslakslaksla1230000@localhost:27017/beta_koding?auto_reconnect"
  when "beta-local"
    "mongodb://beta_koding_user:lkalkslakslaksla1230000@web0.beta.system.aws.koding.com:27017/beta_koding?auto_reconnect"
  when "wan"
    "mongodb://kodingen_user:Cvy3_exwb6JI@184.173.138.98:27017/kodingen?auto_reconnect"
  when "mongohq-dev"
    "mongodb://dev:633939V3R6967W93A@alex.mongohq.com:10065/koding_copy?auto_reconnect"

koding = new Bongo
  mongo   : dbUrl
  models  : require('path').join __dirname, './server/app/bongo/fixedmodels'
  mq      : new Broker {
    host      : "localhost"
    login     : "guest"
    password  : "guest"
    #host      : "web0.beta.system.aws.koding.com"
    #login     : "guest"
    #password  : "x1srTA7!%Vb}$n|S"
  }
koding.on 'auth', (client)->
  koding.models.JVisitor.authenticateClient client, (err, account)->
    koding.handleResponse client.secretName, 'changeLoggedInState', [account]
koding.connect console.log