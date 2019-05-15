import React, { Component } from "react";
import ReactDOM from "react-dom";
import {Search,Dropdown} from "semantic-ui-react";
import {Results} from "./components/Results";

class App extends Component {
	constructor(props) {
		super(props);
		this.state = {
			query: "",
			selectedEngine: "",
			selectedStrategy: "",
			results: []
		}
	}

	onInputChange(event) {
		this.setState({query: event.target.value});
	}

	onSearchBarKeyPress(event) {
		if (event.key === "Enter") {
			this.search();
		}
	}

	search() {
		this.setState({
			results: [
					{
						"title": "The Go Programming Language",
						"url": "https://golang.org/",
						"description": "Go is an open source "
					},
					{
						"title": "GitHub - golang/go: The Go programming languageCached",
						"url": "https://github.com/golang/go",
						"description": "The Go Programming Language.  "
					}
				]
			}
		);
	}

	selectEngine(event, data) {
		this.setState({
			selectedEngine: data.value
		})
	}

	render() {
		let strategies = [
			{key: 1, value: 'cross', text: 'cross'},
			{key: 2, value: 'all', text: 'all'},
			{key: 3, value: 'top', text: 'top'},
		];
		let engines = [
			{key: 1, value: 'all', text: 'all'},
			{key: 2, value: 'google', text: 'google'},
			{key: 3, value: 'bing', text: 'bing'},
			{key: 4, value: 'ask', text: 'ask'},
			{key: 5, value: 'yahoo', text: 'yahoo'},
		];
		return (
			<div className="app-container">
				<Search
					placeholder="Search..."
					showNoResults={false}
					onSearchChange={this.onInputChange}
					onKeyPress={this.onSearchBarKeyPress}
				/>
				<div style={{display: 'block'}}>
					<Dropdown
						placeholder="Strategy"
						fluid
						selection
						options={strategies}
						onChange={this.selectEngine}
					/>
					<Dropdown
						placeholder="Engine"
						fluid
						selection
						options={engines}
					/>
				</div>
				<hr/>
				<Results
					results={this.state.results}
				/>
			</div>
		);
	}
}

ReactDOM.render(<App />, document.getElementById("app"));