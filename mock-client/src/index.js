import { Component } from 'preact';
import 'bulma/css/bulma.css';

import './style';
import Header from './components/Header';
import Broadcaster from './components/Broadcaster';

const createWSHandler = route => ({ route, socket: new WebSocket(`ws://${location.host}/ws/${route}`) });

export default class App extends Component {
	state = {
		broadcasters: []
	}

	getBroadcasters = () => fetch('/api/routes')
		.then(res => res.ok && res.json())
		.then(({ Names }) => Names && this.setState({ broadcasters: Names.map(route => createWSHandler(route)) }));

	addBroadcaster = route => fetch(`/api/${route}`, { method: 'POST' })
		.then(res => res.ok && this.setState(state => ({
			broadcasters:
				state.broadcasters.every(broadcaster => broadcaster.route !== route)
					? [...state.broadcasters, createWSHandler(route)]
					: state.broadcasters
		})));

	removeBroadcaster = route => fetch(`/api/${route}`, { method: 'DELETE' })
		.then(res => res.ok && this.setState(state => {
			state.broadcasters.forEach(broadcaster => broadcaster.route === route && broadcaster.socket.close());
			return { broadcasters: state.broadcasters.filter(broadcaster => broadcaster.route !== route) };
		}));

	componentDidMount() {
		this.getBroadcasters();
	}

	render(_, { broadcasters }) {
		return (
			<div>
				<Header addBroadcaster={this.addBroadcaster} />
				<div class="container is-fluid broadcaster-container">
					{broadcasters.length ?
						<div class="columns is-multiline">
							{broadcasters.map(
								broadcaster => (
									<Broadcaster
										broadcaster={broadcaster}
										removeBroadcaster={this.removeBroadcaster}
									/>
								)
							)}
						</div> :
						<p class="has-text-centered is-size-3 broadcasters-empty">
							No broadcasters
						</p>
					}
				</div>
			</div >
		);
	}
}
