import { Component } from 'preact';
import 'bulma/css/bulma.css';

import './style';
import Header from './components/Header';
import Broadcaster from './components/Broadcaster';

export default class App extends Component {
	state = {
		broadcasters: []
	}

	addBroadcaster = route => fetch(`/api/${route}`, { method: 'POST' })
		.then(res => res.ok && this.setState(state => ({
			broadcasters:
				state.broadcasters.every(broadcaster => broadcaster.route !== route)
					? [...state.broadcasters, { route, socket: new WebSocket(`ws://${location.host}/ws/${route}`) }]
					: state.broadcasters
		})))

	removeBroadcaster = route => fetch(`/api/${route}`, { method: 'DELETE' })
		.then(res => res.ok && this.setState(state => {
			const isToBeRemoved = broadcaster => broadcaster.route === route;
			state.broadcasters.forEach(broadcaster => isToBeRemoved(broadcaster) && broadcaster.socket.close());
			return { broadcasters: state.broadcasters.filter(isToBeRemoved) };
		}))

	render(_, { broadcasters }) {
		return (
			<div>
				<Header addBroadcaster={this.addBroadcaster} />
				<div class="container is-fluid">
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
