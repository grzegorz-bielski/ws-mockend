import './style';
import { Component } from 'preact';

const addRoute = route => fetch('/api/orders', { method: 'POST' });

export default class App extends Component {
	render() {
		return (
			<div>
				<h1>Hello, World!</h1>
				<button onClick={addRoute} />
			</div>
		);
	}
}
