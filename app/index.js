var Generator = require('yeoman-generator');
var path = require('path');
var mkdirp = require('mkdirp')

module.exports = class extends Generator {
  // The name `constructor` is important here
  constructor(args, opts) {
    // Calling the super constructor is important so our generator is correctly set up
    super(args, opts);
  }

  prompting() {

      let currentDirectory = path.parse(this.destinationRoot())

      return this.prompt([{
      type    : 'input',
      name    : 'package_name',
      message : 'Your package name',
      default : currentDirectory.name // Default to current folder name
    },
    {
      type    : 'input',
      name    : 'model_name',
      message : 'What is your model name ?',
    }]).then((answers) => {
        this.props = {}
        this.props.package_name = answers.package_name.replace(' ', '-')
        this.props.model_name = answers.model_name
        this.props.camel_model_name = answers.model_name.charAt(0).toLowerCase() + answers.model_name.slice(1);
    });;
  };

  writing(){
      this.fs.copyTpl(this.templatePath('endpoint.go'), this.destinationPath('./endpoint.go'), 
      { 
        package_name: this.props.package_name,
        model_name: this.props.model_name,
        camel_model_name: this.props.camel_model_name,
      })
      this.fs.copyTpl(this.templatePath('endpoint_test.go'), this.destinationPath('./endpoint_test.go'), 
      { 
        package_name: this.props.package_name,
        model_name: this.props.model_name,
        camel_model_name: this.props.camel_model_name,
      })
      this.fs.copyTpl(this.templatePath('http_transport.go'), this.destinationPath('./http_transport.go'), 
      { 
        package_name: this.props.package_name,
        model_name: this.props.model_name,
        camel_model_name: this.props.camel_model_name,
      })
      this.fs.copyTpl(this.templatePath('http_transport_test.go'), this.destinationPath('./http_transport_test.go'), 
      { 
        package_name: this.props.package_name,
        model_name: this.props.model_name,
        camel_model_name: this.props.camel_model_name,
      })
      this.fs.copyTpl(this.templatePath('model.go'), this.destinationPath('./model.go'), 
      { 
        package_name: this.props.package_name,
        model_name: this.props.model_name,
        camel_model_name: this.props.camel_model_name,
      })
      this.fs.copyTpl(this.templatePath('service.go'), this.destinationPath('./service.go'), 
      { 
        package_name: this.props.package_name,
        model_name: this.props.model_name,
        camel_model_name: this.props.camel_model_name,
      })
      this.fs.copyTpl(this.templatePath('service_test.go'), this.destinationPath('./service_test.go'), 
      { 
        package_name: this.props.package_name,
        model_name: this.props.model_name,
        camel_model_name: this.props.camel_model_name,
      })
      this.fs.copyTpl(this.templatePath('service_tracing.go'), this.destinationPath('./service_tracing.go'), 
      { 
        package_name: this.props.package_name,
        model_name: this.props.model_name,
        camel_model_name: this.props.camel_model_name,
      })
      this.fs.copyTpl(this.templatePath('service_tracing_test.go'), this.destinationPath('./service_tracing_test.go'), 
      { 
        package_name: this.props.package_name,
        model_name: this.props.model_name,
        camel_model_name: this.props.camel_model_name,
      })
      this.fs.copyTpl(this.templatePath('_main.todelete.txt'), this.destinationPath('./_main.todelete.txt'), 
      { 
        package_name: this.props.package_name,
        model_name: this.props.model_name,
        camel_model_name: this.props.camel_model_name,
      })
  }

};