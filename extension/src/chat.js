const {ChatClient} = require('./chat_grpc_web_pb')
const {ConnReq, GroupActionReq, Msg, PubMsgReq} = require('./chat_pb')

function chatNewClient(host) {
	return new ChatClient(host, null, null)
}

function chatConn(client, meta, token, callback) {
	var connReq = new ConnReq()
	connReq.setToken(token)
	var stream = client.conn(connReq, meta)
	stream.on('data', (rsp) => {
		console.log('conn data', rsp)
		var msg = {
			fromId: rsp.getFromId(),
			toId: rsp.getToId(),
			content: rsp.getContent(),
		}
		callback(msg, null)
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

function chatJoinGroup(client, meta, url) {
	var req = new GroupActionReq()
	req.setAction('join')
	req.setGid(url)
	client.groupAction(req, meta, (err, rsp) => {
		if (err) {
			console.error('json group', url, err)
		} else {
			console.log(rsp)
		}
	})
}

function chatPubmsg(client, meta, gid, content) {
	var msg = new Msg()
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

module.exports = {
  chatNewClient: chatNewClient,
  chatConn: chatConn,
  chatPubmsg: chatPubmsg,
  chatJoinGroup: chatJoinGroup,
};
