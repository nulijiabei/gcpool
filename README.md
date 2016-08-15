
一个基于WebSocket(or net.Conn)多连接数据推送核心 ... 模型

-------------

	gcpool "github.com/nulijiabei/go-conn-pool"

	// 节点连接池
	var GO_CONN_POOL *gcpool.Pool
	
	func main() {
		
		// 启动 PCore 服务
		GO_CONN_POOL = gcpool.NewPool()
		GO_CONN_POOL.Register("default")
		GO_CONN_POOL.Start()
		
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
		
		// 保存连接
		GO_CONN_POOL.GetConn("default").Add(device, ws)
		// 断开移除
		defer GO_CONN_POOL.GetConn("default").Del(device)
			
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