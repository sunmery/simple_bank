import { createLazyFileRoute } from '@tanstack/react-router'
import AppBar from '../components/AppBar'

export const Route = createLazyFileRoute('/')({
  component: Index,
})

function Index() {
  return <AppBar />
}
