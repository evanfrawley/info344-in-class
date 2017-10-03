import React, { Component } from 'react';
import 'whatwg-fetch';
import './App.css';


class App extends Component {

    constructor(props) {
        super(props);
        this.state = {
            phrase: "",
            mem: ""
        }
    }

    componentWillMount() {
        setInterval(this.fetchMemResource, 1000)
    }

  fetchNameResource(name) {
      let str = `http://localhost:4000/hello?name=${name}`;
      fetch(str).then((response) => {
          return response.text();
      }).then((text) => {
          this.setState({phrase: text})
      })
  }

  fetchMemResource = () => {
      let str = "http://localhost:4000/meme";
      fetch(str).then((response) => {
          return response.json();
      }).then((json) => {
          this.setState({mem: json["Alloc"]})
      })
  };

  handleSubmit(event) {
      event.preventDefault();
      let name = event.target.name.value;
      this.fetchNameResource(name);
  }

  render() {
    return (
      <div className="App">
        <h1>This is the in class exercise for info344</h1>
        <div>
          <form onSubmit={this.handleSubmit.bind(this)}>
            Name: <input type="text" name="name"/>
            <input type="submit" value={"Submit"}/>
          </form>
        </div>
        <div>{this.state.phrase}</div>
        <div>Used memory: {this.state.mem}</div>
      </div>
    );
  }
}

export default App;
