// Seznam zdrojových kódů v tomto podadresáři
// ------------------------------------------
//
// 01_simple_client.go:
//    Standardní balíček "net".
//    Klient, který přečte ze serveru sekvenci bajtů
//
// 01B_simple_client_headers.go:
//    Standardní balíček "net".
//    Upravený klient, který vytiskne místní i vzdálenou adresu
//
// 02_simple_server.go:
//    Standardní balíček "net".
//    Jednoduchý server posílající jediný bajt klientovi
//
// 02B_simple_server_no_localhost.go:
//    Standardní balíček "net".
//    Úprava adresy v předchozím příkladu
//
// 03_slow_server.go:
//    Standardní balíček "net".
//    Server odpovídající klientovi opožděně
//
// 04_multi_connection_server.go:
//    Standardní balíček "net".
//    Server, který dokáže obsloužit více klientů současně
//
// 05_text_client.go:
//    Standardní balíček "net".
//    Jednoduchý klient akceptující celý textový řádek
//
// 06_text_server.go:
//    Standardní balíček "net".
//    Server posílající klientovi textová data
//
// 06B_better_text_server.go:
//    Standardní balíček "net".
//    Vylepšený server posílající klientovi textová data
//
// 06C_wrong_connection_close.go:
//    Standardní balíček "net".
//    Připojení je nutné ukončit v gorutině, ne mimo ni
//
// 07_lookup.go:
//    Standardní balíček "net".
//    Překlad doménového jména na IP adresy
//
// 08_parse_ip.go:
//    Standardní balíček "net".
//    Parsing IPv4 a IPv6 adresy
//
// 09_ipv4_constructor.go:
//    Standardní balíček "net".
//    Konstruktor adresy typu IPv4
//
// 10_http_get.go:
//    Standardní balíček "net".
//    Použití HTTP metody GET
//
// 11_http_print_headers.go:
//    Standardní balíček "net".
//    Vytištění hlavičky HTTP odpovědi
//
// 12_http_server.go:
//    Standardní balíček "net".
//    Nejjednodušší HTTP server s jediným endpointem
//
// 13_http_server_with_state.go:
//    Standardní balíček "net".
//    HTTP server se stavovou proměnnou
//
// 14_http_server_with_state_mutex.go:
//    Standardní balíček "net".
//    Korektní práce se stavovou proměnnou
//
// 15_file_server.go:
//    Standardní balíček "net".
//    HTTP server vracející statický obsah
//
// 16_custom_server.go:
//    Standardní balíček "net".
//    Kombinace předchozích možností
//
