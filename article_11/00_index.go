// Seznam zdrojových kódů v tomto podadresáři
// ------------------------------------------
//
// 01_simple_client.go:
//    Klient, který přečte ze serveru sekvenci bajtů
//
// 01B_simple_client_headers.go:
//    Upravený klient, který vytiskne místní i vzdálenou adresu
//
// 02_simple_server.go:
//     Jednoduchý server posílající jediný bajt klientovi
//
// 02B_simple_server_no_localhost.go:
//     Úprava adresy v předchozím příkladu
//
// 03_slow_server.go:
//     Server odpovídající klientovi opožděně
//
// 04_multi_connection_server.go:
//     Server, který dokáže obsloužit více klientů současně
//
// 05_text_client.go:
//     	Jednoduchý klient akceptující celý textový řádek
//
// 06_text_server.go:
//     Server posílající klientovi textová data
//
// 06B_better_text_server.go:
//     	Vylepšený server posílající klientovi textová data
//
// 06C_wrong_connection_close.go:
//     Připojení je nutné ukončit v gorutině, ne mimo ni
//
// 07_lookup.go:
//     Překlad doménového jména na IP adresy
//
// 08_parse_ip.go:
//     Parsing IPv4 a IPv6 adresy
//
// 09_ipv4_constructor.go:
//     Konstruktor adresy typu IPv4
//
// 10_http_get.go:
//     Použití HTTP metody GET
//
// 11_http_print_headers.go:
//    Vytištění hlavičky HTTP odpovědi
//
// 12_http_server.go:
//     Nejjednodušší HTTP server s jediným endpointem
//
// 13_http_server_with_state.go:
//     HTTP server se stavovou proměnnou
//
// 14_http_server_with_state_mutex.go:
//     Korektní práce se stavovou proměnnou
//
// 15_file_server.go:
//     HTTP server vracející statický obsah
//
// 16_custom_server.go:
//     Kombinace předchozích možností
//
