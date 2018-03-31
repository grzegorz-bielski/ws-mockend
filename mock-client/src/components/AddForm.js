import { h, Component } from 'preact';

export default class AddForm extends Component {
	handleSubmit = event => {
		event.preventDefault();
		const { value } = event.target.bcName;
		if (value) {
			this.props.addBroadcaster(value);
		}
	};

	render() {
		return (
			<form class="field has-addons" onSubmit={this.handleSubmit}>
				<p class="control">
					<input class="input" type="text" name="bcName" placeholder="Add broadcaster" />
				</p>
				<p class="control">
					<button class="button">
						Add
					</button>
				</p>
			</form>
		);
	}
}