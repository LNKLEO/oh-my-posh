package runtime

import (
	"strings"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

const (
	// see https://docs.microsoft.com/en-us/windows/win32/api/netioapi/ns-netioapi-mib_if_row2

	// InterfaceType
	IF_TYPE_OTHER              IFTYPE = "Other"            // 1
	IF_TYPE_ETHERNET_CSMACD    IFTYPE = "Ethernet/802.3"   // 6
	IF_TYPE_ISO88025_TOKENRING IFTYPE = "Token Ring/802.5" // 9
	IF_TYPE_FDDI               IFTYPE = "FDDI"             // 15
	IF_TYPE_PPP                IFTYPE = "PPP"              // 23
	IF_TYPE_SOFTWARE_LOOPBACK  IFTYPE = "Loopback"         // 24
	IF_TYPE_ATM                IFTYPE = "ATM"              // 37
	IF_TYPE_IEEE80211          IFTYPE = "Wi-Fi/802.11"     // 71
	IF_TYPE_TUNNEL             IFTYPE = "Tunnel"           // 131
	IF_TYPE_IEEE1394           IFTYPE = "FireWire/1394"    // 144
	IF_TYPE_IEEE80216_WMAN     IFTYPE = "WMAN/802.16"      // 237 WiMax
	IF_TYPE_WWANPP             IFTYPE = "WWANPP/GSM"       // 243 GSM
	IF_TYPE_WWANPP2            IFTYPE = "WWANPP/CDMA"      // 244 CDMA
	IF_TYPE_UNKNOWN            IFTYPE = "Unknown"

	// NDISMediaType
	NdisMedium802_3        NDIS_MEDIUM = "802.3"         // 0
	NdisMedium802_5        NDIS_MEDIUM = "802.5"         // 1
	NdisMediumFddi         NDIS_MEDIUM = "FDDI"          // 2
	NdisMediumWan          NDIS_MEDIUM = "WAN"           // 3
	NdisMediumLocalTalk    NDIS_MEDIUM = "LocalTalk"     // 4
	NdisMediumDix          NDIS_MEDIUM = "DIX"           // 5
	NdisMediumArcnetRaw    NDIS_MEDIUM = "ARCNET"        // 6
	NdisMediumArcnet878_2  NDIS_MEDIUM = "ARCNET(878.2)" // 7
	NdisMediumAtm          NDIS_MEDIUM = "ATM"           // 8
	NdisMediumWirelessWan  NDIS_MEDIUM = "WWAN"          // 9
	NdisMediumIrda         NDIS_MEDIUM = "IrDA"          // 10
	NdisMediumBpc          NDIS_MEDIUM = "Broadcast"     // 11
	NdisMediumCoWan        NDIS_MEDIUM = "CO WAN"        // 12
	NdisMedium1394         NDIS_MEDIUM = "1394"          // 13
	NdisMediumInfiniBand   NDIS_MEDIUM = "InfiniBand"    // 14
	NdisMediumTunnel       NDIS_MEDIUM = "Tunnel"        // 15
	NdisMediumNative802_11 NDIS_MEDIUM = "Native 802.11" // 16
	NdisMediumLoopback     NDIS_MEDIUM = "Loopback"      // 17
	NdisMediumWiMax        NDIS_MEDIUM = "WiMax"         // 18
	NdisMediumUnknown      NDIS_MEDIUM = "Unknown"

	// NDISPhysicalMeidaType
	NdisPhysicalMediumUnspecified  NDIS_PHYSICAL_MEDIUM = "Unspecified"   // 0
	NdisPhysicalMediumWirelessLan  NDIS_PHYSICAL_MEDIUM = "Wireless LAN"  // 1
	NdisPhysicalMediumCableModem   NDIS_PHYSICAL_MEDIUM = "Cable Modem"   // 2
	NdisPhysicalMediumPhoneLine    NDIS_PHYSICAL_MEDIUM = "Phone Line"    // 3
	NdisPhysicalMediumPowerLine    NDIS_PHYSICAL_MEDIUM = "Power Line"    // 4
	NdisPhysicalMediumDSL          NDIS_PHYSICAL_MEDIUM = "DSL"           // 5
	NdisPhysicalMediumFibreChannel NDIS_PHYSICAL_MEDIUM = "Fibre Channel" // 6
	NdisPhysicalMedium1394         NDIS_PHYSICAL_MEDIUM = "1394"          // 7
	NdisPhysicalMediumWirelessWan  NDIS_PHYSICAL_MEDIUM = "Wireless WAN"  // 8
	NdisPhysicalMediumNative802_11 NDIS_PHYSICAL_MEDIUM = "Native 802.11" // 9
	NdisPhysicalMediumBluetooth    NDIS_PHYSICAL_MEDIUM = "Bluetooth"     // 10
	NdisPhysicalMediumInfiniband   NDIS_PHYSICAL_MEDIUM = "Infini Band"   // 11
	NdisPhysicalMediumWiMax        NDIS_PHYSICAL_MEDIUM = "WiMax"         // 12
	NdisPhysicalMediumUWB          NDIS_PHYSICAL_MEDIUM = "UWB"           // 13
	NdisPhysicalMedium802_3        NDIS_PHYSICAL_MEDIUM = "802.3"         // 14
	NdisPhysicalMedium802_5        NDIS_PHYSICAL_MEDIUM = "802.5"         // 15
	NdisPhysicalMediumIrda         NDIS_PHYSICAL_MEDIUM = "IrDA"          // 16
	NdisPhysicalMediumWiredWAN     NDIS_PHYSICAL_MEDIUM = "Wired WAN"     // 17
	NdisPhysicalMediumWiredCoWan   NDIS_PHYSICAL_MEDIUM = "Wired CO WAN"  // 18
	NdisPhysicalMediumOther        NDIS_PHYSICAL_MEDIUM = "Other"         // 19
	NdisPhysicalMediumUnknown      NDIS_PHYSICAL_MEDIUM = "Unknown"
)

func (term *Terminal) GetAllNetworkInterfaces() (*[]NetworkInfo, error) {
	var pIFTable2 *MIN_IF_TABLE2
	hGetIfTable2.Call(uintptr(unsafe.Pointer(&pIFTable2)))

	SSIDs, _ := term.GetAllWifiSSID()
	networks := make([]NetworkInfo, 0)

	for i := 0; i < int(pIFTable2.NumEntries); i++ {
		_if := pIFTable2.Table[i]
		_Alias := strings.TrimRight(syscall.UTF16ToString(_if.Alias[:]), "\x00")
		if _if.PhysicalMediumType != 0 && _if.OperStatus == 1 &&
			strings.LastIndex(_Alias, "-") <= 4 &&
			!strings.HasPrefix(_Alias, "Local Area Connection") {
			network := NetworkInfo{}
			network.Alias = _Alias
			network.Interface = strings.TrimRight(syscall.UTF16ToString(_if.Description[:]), "\x00")
			network.TransmitLinkSpeed = _if.TransmitLinkSpeed
			network.ReceiveLinkSpeed = _if.ReceiveLinkSpeed

			switch _if.Type {
			case 1:
				network.InterfaceType = IF_TYPE_OTHER
			case 6:
				network.InterfaceType = IF_TYPE_ETHERNET_CSMACD
			case 9:
				network.InterfaceType = IF_TYPE_ISO88025_TOKENRING
			case 15:
				network.InterfaceType = IF_TYPE_FDDI
			case 23:
				network.InterfaceType = IF_TYPE_PPP
			case 24:
				network.InterfaceType = IF_TYPE_SOFTWARE_LOOPBACK
			case 37:
				network.InterfaceType = IF_TYPE_ATM
			case 71:
				network.InterfaceType = IF_TYPE_IEEE80211
			case 131:
				network.InterfaceType = IF_TYPE_TUNNEL
			case 144:
				network.InterfaceType = IF_TYPE_IEEE1394
			case 237:
				network.InterfaceType = IF_TYPE_IEEE80216_WMAN
			case 243:
				network.InterfaceType = IF_TYPE_WWANPP
			case 244:
				network.InterfaceType = IF_TYPE_WWANPP2
			default:
				network.InterfaceType = IF_TYPE_UNKNOWN
			}

			switch _if.MediaType {
			case 0:
				network.NDISMediaType = NdisMedium802_3
			case 1:
				network.NDISMediaType = NdisMedium802_5
			case 2:
				network.NDISMediaType = NdisMediumFddi
			case 3:
				network.NDISMediaType = NdisMediumWan
			case 4:
				network.NDISMediaType = NdisMediumLocalTalk
			case 5:
				network.NDISMediaType = NdisMediumDix
			case 6:
				network.NDISMediaType = NdisMediumArcnetRaw
			case 7:
				network.NDISMediaType = NdisMediumArcnet878_2
			case 8:
				network.NDISMediaType = NdisMediumAtm
			case 9:
				network.NDISMediaType = NdisMediumWirelessWan
			case 10:
				network.NDISMediaType = NdisMediumIrda
			case 11:
				network.NDISMediaType = NdisMediumBpc
			case 12:
				network.NDISMediaType = NdisMediumCoWan
			case 13:
				network.NDISMediaType = NdisMedium1394
			case 14:
				network.NDISMediaType = NdisMediumInfiniBand
			case 15:
				network.NDISMediaType = NdisMediumTunnel
			case 16:
				network.NDISMediaType = NdisMediumNative802_11
			case 17:
				network.NDISMediaType = NdisMediumLoopback
			case 18:
				network.NDISMediaType = NdisMediumWiMax
			default:
				network.NDISMediaType = NdisMediumUnknown
			}

			switch _if.PhysicalMediumType {
			case 0:
				network.NDISPhysicalMeidaType = NdisPhysicalMediumUnspecified
			case 1:
				network.NDISPhysicalMeidaType = NdisPhysicalMediumWirelessLan
			case 2:
				network.NDISPhysicalMeidaType = NdisPhysicalMediumCableModem
			case 3:
				network.NDISPhysicalMeidaType = NdisPhysicalMediumPhoneLine
			case 4:
				network.NDISPhysicalMeidaType = NdisPhysicalMediumPowerLine
			case 5:
				network.NDISPhysicalMeidaType = NdisPhysicalMediumDSL
			case 6:
				network.NDISPhysicalMeidaType = NdisPhysicalMediumFibreChannel
			case 7:
				network.NDISPhysicalMeidaType = NdisPhysicalMedium1394
			case 8:
				network.NDISPhysicalMeidaType = NdisPhysicalMediumWirelessWan
			case 9:
				network.NDISPhysicalMeidaType = NdisPhysicalMediumNative802_11
			case 10:
				network.NDISPhysicalMeidaType = NdisPhysicalMediumBluetooth
			case 11:
				network.NDISPhysicalMeidaType = NdisPhysicalMediumInfiniband
			case 12:
				network.NDISPhysicalMeidaType = NdisPhysicalMediumWiMax
			case 13:
				network.NDISPhysicalMeidaType = NdisPhysicalMediumUWB
			case 14:
				network.NDISPhysicalMeidaType = NdisPhysicalMedium802_3
			case 15:
				network.NDISPhysicalMeidaType = NdisPhysicalMedium802_5
			case 16:
				network.NDISPhysicalMeidaType = NdisPhysicalMediumIrda
			case 17:
				network.NDISPhysicalMeidaType = NdisPhysicalMediumWiredWAN
			case 18:
				network.NDISPhysicalMeidaType = NdisPhysicalMediumWiredCoWan
			case 19:
				network.NDISPhysicalMeidaType = NdisPhysicalMediumOther
			default:
				network.NDISPhysicalMeidaType = NdisPhysicalMediumUnknown
			}

			if SSID, OK := SSIDs[network.Interface]; OK {
				network.SSID = SSID
			}

			networks = append(networks, network)
		}
	}
	return &networks, nil
}

func (term *Terminal) GetAllWifiSSID() (map[string]string, error) {
	var pdwNegotiatedVersion uint32
	var phClientHandle uint32
	e, _, err := hWlanOpenHandle.Call(uintptr(uint32(2)), uintptr(unsafe.Pointer(nil)), uintptr(unsafe.Pointer(&pdwNegotiatedVersion)), uintptr(unsafe.Pointer(&phClientHandle)))
	if e != 0 {
		return nil, err
	}

	// defer closing handle
	defer func() {
		_, _, _ = hWlanCloseHandle.Call(uintptr(phClientHandle), uintptr(unsafe.Pointer(nil)))
	}()

	ssid := make(map[string]string)
	// list interfaces
	var interfaceList *WLAN_INTERFACE_INFO_LIST
	e, _, err = hWlanEnumInterfaces.Call(uintptr(phClientHandle), uintptr(unsafe.Pointer(nil)), uintptr(unsafe.Pointer(&interfaceList)))
	if e != 0 {
		return nil, err
	}

	numberOfInterfaces := int(interfaceList.dwNumberOfItems)
	for i := 0; i < numberOfInterfaces; i++ {
		infoSize := unsafe.Sizeof(interfaceList.InterfaceInfo[i])
		network := (*WLAN_INTERFACE_INFO)(unsafe.Pointer(uintptr(unsafe.Pointer(&interfaceList.InterfaceInfo[i])) + uintptr(i)*infoSize))
		if network.isState == 1 {
			wifi, _ := term.parseWlanInterface(network, phClientHandle)
			ssid[wifi.Interface] = wifi.SSID
		}
	}
	return ssid, nil
}

const (
	FHSS   WifiType = "FHSS"
	DSSS   WifiType = "DSSS"
	IR     WifiType = "IR"
	A      WifiType = "802.11a"
	HRDSSS WifiType = "HRDSSS"
	G      WifiType = "802.11g"
	N      WifiType = "802.11n"
	AC     WifiType = "802.11ac"

	Infrastructure WifiType = "Infrastructure"
	Independent    WifiType = "Independent"
	Any            WifiType = "Any"

	OpenSystem WifiType = "802.11 Open System"
	SharedKey  WifiType = "802.11 Shared Key"
	WPA        WifiType = "WPA"
	WPAPSK     WifiType = "WPA PSK"
	WPANone    WifiType = "WPA NONE"
	WPA2       WifiType = "WPA2"
	WPA2PSK    WifiType = "WPA2 PSK"
	Disabled   WifiType = "disabled"

	None   WifiType = "None"
	WEP40  WifiType = "WEP40"
	TKIP   WifiType = "TKIP"
	CCMP   WifiType = "CCMP"
	WEP104 WifiType = "WEP104"
	WEP    WifiType = "WEP"
)

func (term *Terminal) parseWlanInterface(network *WLAN_INTERFACE_INFO, clientHandle uint32) (*WifiInfo, error) {
	info := WifiInfo{}
	info.Interface = strings.TrimRight(string(utf16.Decode(network.strInterfaceDescription[:])), "\x00")

	// Query wifi connection state
	var dataSize uint32
	var wlanAttr *WLAN_CONNECTION_ATTRIBUTES
	e, _, err := hWlanQueryInterface.Call(uintptr(clientHandle),
		uintptr(unsafe.Pointer(&network.InterfaceGuid)),
		uintptr(7), // wlan_intf_opcode_current_connection
		uintptr(unsafe.Pointer(nil)),
		uintptr(unsafe.Pointer(&dataSize)),
		uintptr(unsafe.Pointer(&wlanAttr)),
		uintptr(unsafe.Pointer(nil)))
	if e != 0 {
		return &info, err
	}

	// SSID
	ssid := wlanAttr.wlanAssociationAttributes.dot11Ssid
	if ssid.uSSIDLength > 0 {
		info.SSID = string(ssid.ucSSID[0:ssid.uSSIDLength])
	}

	// see https://docs.microsoft.com/en-us/windows/win32/nativewifi/dot11-phy-type
	switch wlanAttr.wlanAssociationAttributes.dot11PhyType {
	case 1:
		info.PhysType = FHSS
	case 2:
		info.PhysType = DSSS
	case 3:
		info.PhysType = IR
	case 4:
		info.PhysType = A
	case 5:
		info.PhysType = HRDSSS
	case 6:
		info.PhysType = G
	case 7:
		info.PhysType = N
	case 8:
		info.PhysType = AC
	default:
		info.PhysType = UNKNOWN
	}

	// see https://docs.microsoft.com/en-us/windows/win32/nativewifi/dot11-bss-type
	switch wlanAttr.wlanAssociationAttributes.dot11BssType {
	case 1:
		info.RadioType = Infrastructure
	case 2:
		info.RadioType = Independent
	default:
		info.RadioType = Any
	}

	info.Signal = int(wlanAttr.wlanAssociationAttributes.wlanSignalQuality)
	info.TransmitRate = int(wlanAttr.wlanAssociationAttributes.ulTxRate) / 1024
	info.ReceiveRate = int(wlanAttr.wlanAssociationAttributes.ulRxRate) / 1024

	// Query wifi channel
	dataSize = 0
	var channel *uint32
	e, _, err = hWlanQueryInterface.Call(uintptr(clientHandle),
		uintptr(unsafe.Pointer(&network.InterfaceGuid)),
		uintptr(8), // wlan_intf_opcode_channel_number
		uintptr(unsafe.Pointer(nil)),
		uintptr(unsafe.Pointer(&dataSize)),
		uintptr(unsafe.Pointer(&channel)),
		uintptr(unsafe.Pointer(nil)))
	if e != 0 {
		return &info, err
	}
	info.Channel = int(*channel)

	if wlanAttr.wlanSecurityAttributes.bSecurityEnabled <= 0 {
		info.Authentication = Disabled
		return &info, nil
	}

	// see https://docs.microsoft.com/en-us/windows/win32/nativewifi/dot11-auth-algorithm
	switch wlanAttr.wlanSecurityAttributes.dot11AuthAlgorithm {
	case 1:
		info.Authentication = OpenSystem
	case 2:
		info.Authentication = SharedKey
	case 3:
		info.Authentication = WPA
	case 4:
		info.Authentication = WPAPSK
	case 5:
		info.Authentication = WPANone
	case 6:
		info.Authentication = WPA2
	case 7:
		info.Authentication = WPA2PSK
	default:
		info.Authentication = UNKNOWN
	}

	// see https://docs.microsoft.com/en-us/windows/win32/nativewifi/dot11-cipher-algorithm
	switch wlanAttr.wlanSecurityAttributes.dot11CipherAlgorithm {
	case 0:
		info.Cipher = None
	case 0x1:
		info.Cipher = WEP40
	case 0x2:
		info.Cipher = TKIP
	case 0x4:
		info.Cipher = CCMP
	case 0x5:
		info.Cipher = WEP104
	case 0x100:
		info.Cipher = WPA
	case 0x101:
		info.Cipher = WEP
	default:
		info.Cipher = UNKNOWN
	}

	return &info, nil
}
