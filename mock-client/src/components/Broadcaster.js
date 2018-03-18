import { h } from 'preact';

const Broadcaster = ({ route }) => (
	<div class="column is-one-quarter">
		<article class="message">
			<div class="message-header">
				<h3>/api/{route}</h3>
				<button class="delete" aria-label="delete" />
			</div>
			<div class="message-body">
				lorem ipsum
			</div>
		</article>
	</div>
);

export default Broadcaster;