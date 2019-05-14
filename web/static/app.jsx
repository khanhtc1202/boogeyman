import React, { Component } from "react";
import ReactDOM from "react-dom";

class App extends Component {
	constructor(props) {
		super(props);
		this.state = {
			query: "sample",
			selectedEngine: "",
			selectedStrategy: ""
		}
	}

	render() {
		return (
			<div className="app-container">
				<p>{this.state.query}</p>
			</div>
		);
	}
}

ReactDOM.render(<App />, document.getElementById("app"));