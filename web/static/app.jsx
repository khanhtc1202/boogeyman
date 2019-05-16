import React, { Component } from "react";
import ReactDOM from "react-dom";
import {Search,Dropdown,Item} from "semantic-ui-react";

class App extends Component {
	constructor(props) {
		super(props);
		this.state = {
			query: "",
			selectedEngine: "",
			selectedStrategy: "",
			results: []
		};
		this.onInputChange = this.onInputChange.bind(this);
		this.onSearchBarKeyPress = this.onSearchBarKeyPress.bind(this);
		this.search = this.search.bind(this);
		this.selectEngine = this.selectEngine.bind(this);
		this.selectStrategy = this.selectStrategy.bind(this);
	}

	onInputChange(event, data) {
		this.setState({query: data.value});
	}

	onSearchBarKeyPress(event) {
		if (event.key === "Enter") {
			this.search();
		}
	}

	search() {
		fetch("http://localhost:3000/search"
			+ "?q=" + this.state.query
			+ "&s=" + this.state.selectedStrategy
			+ "&e=" + this.state.selectedEngine)
			.then(response => response.json())
			.then(data => {
				if (data.results !== null) {
					this.setState({results: data.results});
				}
			})
			.catch(() => {});
	}

	selectEngine(event, data) {
		this.setState({
			selectedEngine: data.value
		})
	}

	selectStrategy(event, data) {
		this.setState({
			selectedStrategy: data.value
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
		let groupedItems = this.state.results.map((item, index) => {
			return (
				<Item>
					<Item.Content>
						<Item.Header>{index+1}. {item.title}</Item.Header>
						<Item.Extra as='a' href={item.url} target='_blank'>{item.url}</Item.Extra>
						<Item.Description>{item.description}</Item.Description>
					</Item.Content>
				</Item>
			);
		});
		return (
			<div className="app-container">
				<div style={{display: "inline-block"}}>
					<Dropdown
						placeholder="Strategy" fluid selection
						options={strategies}
						onChange={this.selectStrategy}
					/>
				</div>
				<div style={{display: 'inline-block'}}>
					<Dropdown
						placeholder="Engine" fluid selection
						options={engines}
						onChange={this.selectEngine}
					/>
				</div>
				<div style={{display: 'inline-block'}}>
					<Search
						placeholder="Search..."
						showNoResults={false}
						onSearchChange={this.onInputChange}
						onKeyPress={this.onSearchBarKeyPress}
					/>
				</div>
				<hr/>
				<Item.Group divided>
					{groupedItems}
				</Item.Group>
			</div>
		);
	}
}

ReactDOM.render(<App />, document.getElementById("app"));