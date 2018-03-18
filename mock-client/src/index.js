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
			broadcasters: [...state.broadcasters, { route }]
		})))

	render(_, { broadcasters }) {
		return (
			<div>
				<Header addBroadcaster={this.addBroadcaster} />
				<div class="container is-fluid">
					<div class="columns is-multiline">{broadcasters ? broadcasters.map(
						({ route }) => <Broadcaster route={route} />
					) : <p class="title">No broadcasters</p>
					}</div>
				</div>
			</div>
		);
	}
}
