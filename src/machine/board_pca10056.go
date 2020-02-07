// +build pca10056

package machine

const HasLowFrequencyCrystal = true

// LEDs on the pca10056
const (
	LED  Pin = LED1
	LED1 Pin = 13
	LED2 Pin = 14
	LED3 Pin = 15
	LED4 Pin = 16
)

// Buttons on the pca10056
const (
	BUTTON  Pin = BUTTON1
	BUTTON1 Pin = 11
	BUTTON2 Pin = 12
	BUTTON3 Pin = 24
	BUTTON4 Pin = 25
)

// UART pins
const (
	UART_TX_PIN Pin = 6
	UART_RX_PIN Pin = 8
)

// UART0 is the NRF UART and UART1 is the nRF52840's USB device
var (
	UART0 = NRF_UART0
	UART1 = USB
)

// ADC pins
const (
	ADC0 Pin = 3
	ADC1 Pin = 4
	ADC2 Pin = 28
	ADC3 Pin = 29
	ADC4 Pin = 30
	ADC5 Pin = 31
)

// I2C pins
const (
	SDA_PIN Pin = 26 // P0.26
	SCL_PIN Pin = 27 // P0.27
)

// SPI pins
const (
	SPI0_SCK_PIN  Pin = 47 // P1.15
	SPI0_MOSI_PIN Pin = 45 // P1.13
	SPI0_MISO_PIN Pin = 46 // P1.14
)
