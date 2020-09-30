package main

import (
	"debug/pe"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %v", err)
	}
}

func main() {
	f, err := os.Open(`C:\Users\Leonty\AppData\Roaming\Telegram Desktop\Telegram.exe`)
	check(err)
	pefile, err := pe.NewFile(f)
	check(err)
	defer f.Close()
	defer pefile.Close()
	dosHeader := make([]byte, 96)
	sizeOffset := make([]byte, 4)
	// Dec to Ascii (searching for MZ)
	_, err = f.Read(dosHeader)
	check(err)
	fmt.Println("[-----DOS Header / Stub-----]")
	fmt.Printf("[+] Magic Value: %s%s\n", string(dosHeader[0]), string(dosHeader[1]))
	// Validate PE+0+0 (Valid PE format)
	pe_sig_offset := int64(binary.LittleEndian.Uint32(dosHeader[0x3c:]))
	f.ReadAt(sizeOffset[:], pe_sig_offset)
	fmt.Println("[-----Signature Header-----]")
	fmt.Printf("[+] LFANEW Value: %s\n", string(sizeOffset))
	/* OUTPUT
	   [-----DOS Header / Stub-----]
	   [+] Magic Value: MZ
	   [-----Signature Header-----]
	   [+] LFANEW Value: PE
	*/
	// Create the reader and read COFF Header
	sr := io.NewSectionReader(f, 0, 1<<63-1)
	_, err = sr.Seek(pe_sig_offset+4, os.SEEK_SET)
	check(err)
	binary.Read(sr, binary.LittleEndian, &pefile.FileHeader)
	// Print File Header
	fmt.Println("[-----COFF File Header-----]")
	fmt.Printf("[+] Machine Architecture: %#x\n", pefile.FileHeader.Machine)
	fmt.Printf("[+] Number of Sections: %#x\n",
		pefile.FileHeader.NumberOfSections)
	fmt.Printf("[+] Size of Optional Header: %#x\n",
		pefile.FileHeader.SizeOfOptionalHeader)
	// Print section names
	fmt.Println("[-----Section Offsets-----]")
	fmt.Printf("[+] Number of Sections Field Offset: %#x\n", pe_sig_offset+6)
	// this is the end of the Signature header (0x7c) + coff (20bytes) + oh32 (224bytes)
	fmt.Printf("[+] Section Table Offset: %#x\n", pe_sig_offset+0xF8)
	/* OUTPUT
	[-----COFF File Header-----]
	[+] Machine Architecture: 0x14c
	[+] Number of Sections: 0x8
	[+] Size of Optional Header: 0xe0
	[-----Section Offsets-----]
	[+] Number of Sections Field Offset: 0x15e
	[+] Section Table Offset: 0x250
	*/
	// Get size of OptionalHeader
	var sizeofOptionalHeader32 = uint16(binary.Size(pe.OptionalHeader32{}))
	var sizeofOptionalHeader64 = uint16(binary.Size(pe.OptionalHeader64{}))
	var oh32 pe.OptionalHeader32
	var oh64 pe.OptionalHeader64
	// Read OptionalHeader
	switch pefile.FileHeader.SizeOfOptionalHeader {
	case sizeofOptionalHeader32:
		fmt.Printf("[+] 32 bit header\n")
		binary.Read(sr, binary.LittleEndian, &oh32)
	case sizeofOptionalHeader64:
		fmt.Printf("[+] 64 bit header\n")
		binary.Read(sr, binary.LittleEndian, &oh64)
	}
	// Print Optional Header
	fmt.Println("[-----Optional Header-----]")
	fmt.Printf("[+] Entry Point: %#x\n", oh32.AddressOfEntryPoint)
	fmt.Printf("[+] ImageBase: %#x\n", oh32.ImageBase)
	fmt.Printf("[+] Size of Image: %#x\n", oh32.SizeOfImage)
	fmt.Printf("[+] Sections Alignment: %#x\n", oh32.SectionAlignment)
	fmt.Printf("[+] File Alignment: %#x\n", oh32.FileAlignment)
	fmt.Printf("[+] Characteristics: %#x\n", pefile.FileHeader.Characteristics)
	fmt.Printf("[+] Size of Headers: %#x\n", oh32.SizeOfHeaders)
	fmt.Printf("[+] Checksum: %#x\n", oh32.CheckSum)
	fmt.Printf("[+] Machine: %#x\n", pefile.FileHeader.Machine)
	fmt.Printf("[+] Subsystem: %#x\n", oh32.Subsystem)
	fmt.Printf("[+] DLLCharacteristics: %#x\n", oh32.DllCharacteristics)
	/* OUTPUT
	[-----Optional Header-----]
	[+] Entry Point: 0x169e682
	[+] ImageBase: 0x400000
	[+] Size of Image: 0x3172000
	[+] Sections Alignment: 0x1000
	[+] File Alignment: 0x200
	[+] Characteristics: 0x102
	[+] Size of Headers: 0x400
	[+] Checksum: 0x2e41078
	[+] Machine: 0x14c
	[+] Subsystem: 0x2
	[+] DLLCharacteristics: 0x8140
	*/
	// Print Data Directory
	fmt.Println("[-----Data Directory-----]")
	var winnt_datadirs = []string{
		"IMAGE_DIRECTORY_ENTRY_EXPORT",
		"IMAGE_DIRECTORY_ENTRY_IMPORT",
		"IMAGE_DIRECTORY_ENTRY_RESOURCE",
		"IMAGE_DIRECTORY_ENTRY_EXCEPTION",
		"IMAGE_DIRECTORY_ENTRY_SECURITY",
		"IMAGE_DIRECTORY_ENTRY_BASERELOC",
		"IMAGE_DIRECTORY_ENTRY_DEBUG",
		"IMAGE_DIRECTORY_ENTRY_COPYRIGHT",
		"IMAGE_DIRECTORY_ENTRY_GLOBALPTR",
		"IMAGE_DIRECTORY_ENTRY_TLS",
		"IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG",
		"IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT",
		"IMAGE_DIRECTORY_ENTRY_IAT",
		"IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT",
		"IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR",
		"IMAGE_NUMBEROF_DIRECTORY_ENTRIES",
	}
	for idx, directory := range oh32.DataDirectory {
		fmt.Printf("[!] Data Directory: %s\n", winnt_datadirs[idx])
		fmt.Printf("[+] Image Virtual Address: %#x\n", directory.VirtualAddress)
		fmt.Printf("[+] Image Size: %#x\n", directory.Size)
	}
	/* OUTPUT
	   [-----Data Directory-----]
	   [!] Data Directory: IMAGE_DIRECTORY_ENTRY_EXPORT
	   [+] Image Virtual Address: 0x2a7b6b0
	   [+] Image Size: 0x116c
	   [!] Data Directory: IMAGE_DIRECTORY_ENTRY_IMPORT
	   [+] Image Virtual Address: 0x2a7c81c
	   [+] Image Size: 0x12c
	   --snip--
	*/
	fmt.Println("[-----Section Table-----]")
	for _, section := range pefile.Sections { ❶
	fmt.Println("[+] --------------------")
	fmt.Printf("[+] Section Name: %s\n", section.Name)
	fmt.Printf("[+] Section Characteristics: %#x\n", section.Characteristics)
	fmt.Printf("[+] Section Virtual Size: %#x\n", section.VirtualSize)
	fmt.Printf("[+] Section Virtual Offset: %#x\n", section.VirtualAddress)
	fmt.Printf("[+] Section Raw Size: %#x\n", section.Size)
	fmt.Printf("[+] Section Raw Offset to Data: %#x\n", section.Offset)
	fmt.Printf("[+] Section Append Offset (Next Section): %#x\n",
	section.Offset+section.Size)
	}
	/* OUTPUT
	[-----Section Table-----]
	[+] --------------------
	[+] Section Name: .text
	[+] Section Characteristics: 0x60000020 
  [+] Section Virtual Size: 0x1853dd0 
  [+] Section Virtual Offset: 0x1000 
  [+] Section Raw Size: 0x1853e00 
  [+] Section Raw Offset to Data: 0x400 
  [+] Section Append Offset (Next Section): 0x1854200 
  [+] --------------------
  [+] Section Name: .rodata
  [+] Section Characteristics: 0x60000020
  [+] Section Virtual Size: 0x1b00
  [+] Section Virtual Offset: 0x1855000
  [+] Section Raw Size: 0x1c00
  [+] Section Raw Offset to Data: 0x1854200
  [+] Section Append Offset (Next Section): 0x1855e00
--snip--
*/
}
