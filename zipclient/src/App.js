import React, {Component} from 'react';
import 'whatwg-fetch';
import './App.css';

import {getZipsForCityName} from './services/zipService';

class App extends Component {

    constructor(props) {
        super(props);
        this.state = {
            zips: []
        }
    }

    componentWillMount() {

    }


    handleSubmit(event) {
        event.preventDefault();
        let name = event.target.name.value;
        getZipsForCityName(name).then((json) => {
            console.log("json", json);
            this.setState({zips: json})
        })
    }

    render() {
        let zips = this.state.zips.map((zip) => {
            let zipCode = zip["Code"];
            return <option value={zipCode}>{zipCode}</option>
        });
        let stateSet = new Set();

        this.state.zips.forEach((zip) => {
            let state = zip["State"];
            stateSet.add(state);
        });
        let stateArray = [];
        stateSet.forEach((state) => {
            stateArray.push(<option value={state}>{state}</option>)
        });
        return (
            <div className="App">
                <h1>This is the in class exercise for info344</h1>
                <div>
                    <form onSubmit={this.handleSubmit.bind(this)}>
                        City Name: <input type="text" name="name"/>
                        <input type="submit" value={"Submit"}/>
                    </form>
                </div>
                <select>
                    <option selected value="base">Please Select ZIP</option>
                    {zips}
                </select>
                <select>
                    <option selected value="base">Please Select State</option>
                    {stateArray}
                </select>
            </div>
        );
    }
}

export default App;
