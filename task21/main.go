package main

func main() {
	client := &client{}
	mac := &mac{}
	client.insertSquareUsbInComputer(mac)
	windowsMachine := &windows{}
	// создание адаптера
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}
	// передача адаптера в метод ожидающий адаптированный интерфейс
	client.insertSquareUsbInComputer(windowsMachineAdapter)
}
