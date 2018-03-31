import { h, Component } from 'preact';

const placeholder = 'Add broadcaster';
export default class AddForm extends Component {
	state = {
		value: null
	}

	handleChange = event => this.setState({ value: event.target.value });
	handleSubmit = event => {
		const { value } = this.state;
		event.preventDefault();
		if (value && value !== placeholder) {
			this.props.addBroadcaster(value);
		}
	};

	render() {
		return (
			<form class="field has-addons" onSubmit={this.handleSubmit}>
				<p class="control">
					<input
						class="input"
						type="text"
						value={this.state.value}
						onChange={this.handleChange}
						placeholder={placeholder}
					/>
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