import { h } from 'preact';

import AddForm from './AddForm';

const Header = ({ addBroadcaster }) => (
	<header class="hero is-warning">
		<div class="hero-body">
			<div class="level">
				<div class="level-left">
					<div class="level-item">
						<h1 class="title">
							WS Mockend
						</h1>
					</div>
				</div>
				<div class="level-right">
					<div class="level-item">
						<AddForm addBroadcaster={addBroadcaster} />
					</div>
				</div>
			</div>
		</div>
	</header>
);

export default Header;
