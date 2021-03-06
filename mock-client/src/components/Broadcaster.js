import { h, Component } from 'preact';

class Broadcaster extends Component {
	state = {
		value: ''
	}
	handleRemove = () => this.props.removeBroadcaster(this.props.broadcaster.route);
	handleChange = event => this.setState({ value: event.target.value });
	handleSubmit = event => {
		const { broadcaster: { socket } } = this.props;
		event.preventDefault();
		if (socket.readyState === 1) {
			socket.send(this.state.value);
		}
	}

	render({ broadcaster }) {
		return (
			<div class="column is-half">
				<article class="message broadcaster">
					<div class="message-header is-primary">
						<h3>/ws/{broadcaster.route}</h3>
						<button class="delete" aria-label="delete" onClick={this.handleRemove} />
					</div>
					<div class="message-body">
						<form onSubmit={this.handleSubmit} >
							<div class="control">
								<textarea
									class="textarea"
									rows="10"
									value={this.state.value}
									onChange={this.handleChange}
								/>
							</div>
							<div class="control">
								<button class="button broadcast-btn is-warning">
									Broadcast
								</button>
							</div>
						</form>
					</div>
				</article>
			</div>
		);
	}
}

export default Broadcaster;