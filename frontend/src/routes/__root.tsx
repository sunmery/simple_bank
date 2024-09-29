import { createRootRoute, Link, Outlet } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'

export const Route = createRootRoute({
	component: () => (
		<>
			<div >
				<Link to="/" >
					Index
				</Link>{' '}
				<Link
				to={'/login'}>
					Login
				</Link>
			</div>
			<Outlet />
			<TanStackRouterDevtools />
		</>
	),
})
