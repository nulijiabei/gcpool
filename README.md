<a href="https://godoc.org/github.com/nulijiabei/gcpool"><img src="https://godoc.org/github.com/nulijiabei/gcpool?status.svg" alt="GoDoc"></a>

Go WebSocket 连接池（及高效的数据推送） + 后续将支持 net.Conn

-------------

当注册连接池时会自动创建对应的流池

向连接池中添加连接时需要为该链接定义唯一名称 ... 

如需向连接池中的连接写入数据时, 只需向流中唯一名称写入数据 ...

-------------

使用场景：

	比如：微信消息推送 ...
	
	初始化 ...
	// 注册微信用户连接池
	GO_CONN_POOL.Register("WeiXinUser")
	...
	
	当一个微信用户连接时 ...
	// 获取微信用户ID
	id := ws.Request().FormValue("wxid")
	// 保存微信连接到连接池
	GO_CONN_POOL.GetConn("WeiXinUser").Add(id, ws)
	// 当微信连接断开则移除
	defer GO_CONN_POOL.GetConn("WeiXinUser").Del(id)	
	// 读取微信用户发来的数据
	r := bufio.NewReader(ws)
	for {
		// (JSON)
		data, err := r.ReadBytes('\n')
		if err != nil {
			break
		}
		// 比如：用户数据包含目标ID及内容
		wxid = <- (JSON)
		data = <- (JSON)
		// 将数据写入到目标用户连接内 ...
		GO_CONN_POOL.GetStream("WeiXinUser").Add(wxid, data)
	}
	...
	
-------------

	import (
		gcpool "github.com/nulijiabei/gcpool"
	)

	// 节点连接池
	var GO_CONN_POOL *gcpool.Pool
	
	func main() {
		
		// ------------- 关键 -------------------- // 
		GO_CONN_POOL = gcpool.NewPool() // 创建对象
		GO_CONN_POOL.Register("default") // 注册连接池
		GO_CONN_POOL.Start() // 启动服务
		// ------------- -- -------------------- // 
		
		// 创建 HTTP + WebSocket 服务
		http.Handle("/hello", websocket.Handler(HelloHandler))
		// 启动服务 ...
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Panic(err)
		}
		
	}
	
	// WS Handler
	func HelloHandler(ws *websocket.Conn) {
		
		// WS 获取请求参数 ... 
		if err := ws.Request().ParseForm(); err != nil {
			return
		}
	
		// 作为唯一标识符
		id := ws.Request().FormValue("id")
		
		// ------------- 关键 -------------------- // 
		// 保存连接
		GO_CONN_POOL.GetConn("default").Add(id, ws)
		// 断开移除
		defer GO_CONN_POOL.GetConn("default").Del(id)
		// ------------- -- -------------------- // 
					
		// 读取 ... 阻塞
		r := bufio.NewReader(ws)
		for {
			// 按行读取 ... (JSON)
			data, err := r.ReadBytes('\n')
			if err != nil {
				if err != io.EOF {
					// 异常 ...
				}
				// 异常时跳出循环 ... 断开连接 ...
				break
			}
		}
		
	}

---


	注意：完成以上部分 ... 此时可以使用 
	
	// 添加连接
	GO_CONN_POOL.GetConn("default").Add(device, ws)
	// 移除连接
	GO_CONN_POOL.GetConn("default").Del(device)
	// 推送数据到连接
	GO_CONN_POOL.GetStream("default").Add(device, data)
	// 或者你可以创建更多的 Conn 或 Stream 交叉使用 ...
	
---
