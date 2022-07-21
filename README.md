# INI parser
#### Using Go to deal with INI


## Features

#### 1. Load from files and strings.
#### 2. Remove comments, empty lines, white spaces and tabs
#### 3. Get sections names as a slice. 
#### 4. Get parsed data as a map.
#### 5. Get a value of key in section
#### 5. Change values and add keys and sections.
#### 6. Get parsed data as a string.
#### 7. Export parsed data to INI file.

## Usage

#### 1. Create an object of Parser 
`parser := Parser{}`
#### 2. Get INI input
>  - From string
> `parser.LoadFromString(iniTemplate)` 
> - From file
> `parser.LoadFromFile("/fileName")`
#### 3. Use methods
>  - Get sections names as a slice. 
> `parser.GetSectionNames()`
> - Get parsed data as a map.
> `parser.parser1.GetSections()`
> - Get a value of key in section
> `parser.Get("EXAMPLE","example")`
> - Change values and add keys and sections.
> `parser.Set("EXAMPLE","example","{test}")`
> - Get parsed data as a string.
> `parser.String()`
> - Export parsed data to INI file
> `SaveToFile("fileName",outINI_string)`





