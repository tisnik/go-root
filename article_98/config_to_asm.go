func main() {
        configuration := ConfigStruct{}
        GetStorageConfigurationByValue(configuration)
        GetStorageConfigurationByReference(&amp;configuration)
        configuration.GetStorageConfigurationByValue()
        configuration.GetStorageConfigurationByReference()
}
