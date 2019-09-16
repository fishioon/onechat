const {ChatClient} = require('./chat_grpc_web_pb')
const {ConnReq, GroupActionReq, Msg, PubMsgReq} = require('./chat_pb')

var myTabs = []
var lastTabid = 0
var meta = {}
var uid = 'anyone'
var client = new ChatClient('https://onechat.fishioon.com:1443', null, null)

chatInit()

chrome.tabs.onUpdated.addListener(function (tabId, _, tab) {
	console.log('tabid', tabId)
	lastTabid = tabId;
	if (!myTabs[tabId]) {
		myTabs[tabId] = {url: tab.url, gid: 'init'}
		chatJoinGroup(client, meta, tab.url, gid => {
			console.log('join group success', tab.url, gid)
			myTabs[tabId] = {url: tab.url, gid: gid}
			chatPubmsg(client, meta, gid, 'hello')
		})
	}
})

// ===============
function showmsg(tabid, msg) {
	chrome.tabs.sendMessage(tabid, {
		type: 'danmu',
		message: {user: {name: msg.fromId}, content: {text: msg.content, image: ''}}
	});
}

function chatInit() {
	chrome.storage.local.get('onetoken', res => {
		var token = ''
		if (res && res.onetoken) {
			token = res.onetoken
		} else {
			token = randomToken()
			chrome.storage.local.set({'onetoken': token})
		}
		console.log('token', token)
		uid = token.split('-')[0]
		meta = {'Authorization': 'Bearer ' + token}
		chatConnect(client, token)
	})
}

function chatConnect(client, token) {
	var connReq = new ConnReq()
	connReq.setToken(token)
	var stream = client.conn(connReq, {})
	stream.on('data', (rsp) => {
		console.log('conn data', rsp)
		var msg = {
			fromId: rsp.getFromId(),
			toId: rsp.getToId(),
			content: rsp.getContent(),
		}
		showmsg(lastTabid, msg)
	})
	stream.on('status', status => {
		console.log('status', status)
	})
	stream.on('error', err => {
		console.log('error', err)
	})
	stream.on('end', () => {
		console.log('end', 'stream end signal received')
	})
}

function chatJoinGroup(client, meta, url, callback) {
	var req = new GroupActionReq()
	req.setAction('join')
	req.setGid(url)
	client.groupAction(req, meta, (err, rsp) => {
		if (err) {
			console.log('join group error', url, err)
		} else {
			callback(rsp.getGid())
		}
	})
}

function chatPubmsg(client, meta, gid, content) {
	var msg = new Msg()
	msg.setFromId(uid)
	msg.setToId(gid)
	msg.setContent(content)
	var req = new PubMsgReq()
	req.setMsg(msg)
	client.pubMsg(req, meta, (err, rsp) => {
		if (err) {
			console.error('pub msg', msg, err)
		} else {
			console.log(rsp)
		}
	})
}

function randomToken() {
	const randomNames = ['mouse', 'ox', 'tiger', 'rabbit', 'dragon', 'snake', 'horse', 'sheep', 'monkey', 'rooster', 'dog', 'pig']
	var id = makeid(20)
	var n = hashCode(id)
	return randomNames[n % randomNames.length] + '-' + id
}

function makeid(length) {
	var result = ''
	var characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
	var charactersLength = characters.length
	for (var i = 0; i < length; i++) {
		result += characters.charAt(Math.floor(Math.random() * charactersLength))
	}
	return result
}

function hashCode(s) {
	var h = 0, l = s.length, i = 0;
	if (l > 0)
		while (i < l)
			h = (h << 5) - h + s.charCodeAt(i++) | 0;
	return h;
}
