import { h, Component } from 'preact';

class Broadcaster extends Component {
	state = {
		content: ''
	}

	render({ broadcaster }) {
		return (
			<div class="column is-one-quarter">
				<article class="message">
					<div class="message-header">
						<h3>/api/{broadcaster.route}</h3>
						<button class="delete" aria-label="delete" />
					</div>
					<div class="message-body">
						<textarea class="textarea" rows="10" value={this.state.content} />
					</div>
				</article>
			</div>
		);
	}
}

export default Broadcaster;