# yeoman generator for microservice in go

I want to generate the scaffolding of my gokit folder with all files and already unit tested.

## Install

Clone this repository (git clone)

If yeoman is not installed please
```bash
npm install -g yo
```

Open the solution
```bash
npm install
sudo npm link
```

now that we are ready to test.

```bash
mkdir [ProjectName]
cd [ProjectName]
yo gokit-folder
```

You can test it :

```bash
go test ./... -tags=unit -v
```

In order to run it, you have to create the endpoints in your main.go
To help you a _main.todelete.txt is generated.
Copy/paste and adpat maybe a little bit.

```bash
go run main.go
```

Then you'll be able to test the endpoints :


```bash
GET /model_Plural_name_choosen/1
{
    "model_nameID" : 1
}
```

```bash
POST /model_Plural_name_choosen/
body {
    "model_nameID" : 1
}

201 Location : "/model_Plural_name_choosen/1
```