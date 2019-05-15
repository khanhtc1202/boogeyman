import React, { Component } from "react";
import ReactDOM from "react-dom";
import { Result } from "components/Result";
import { SearchBar } from "components/SearchBar";

class App extends Component {
	constructor(props) {
		super(props);
		this.state = {
			results: []
		}
	}

	render() {
		return (
			<div className="app-container">
				<SearchBar query={()=>{}}/>
				<hr/>
				<Result
					title="sample"
					link="https://khanhtc.me"
					description="sample text"
				/>
			</div>
		);
	}
}

ReactDOM.render(<App />, document.getElementById("app"));