import { createBrowserRouter } from 'react-router-dom'
import Login from '../pages/Login'
import Layout from '../layouts/Layout'
import Memory from '../pages/Memory'
import Discover from '../pages/Discover'
import VoicePlan from '../pages/VoicePlan'
import Guide from '../pages/Guide'
import Profile from '../pages/Profile'

const router = createBrowserRouter([
  {
    path: '/login',
    element: <Login />,
  },
  {
    path: '/',
    element: <Layout />,
    children: [
      {
        path: 'memory',
        element: <Memory />,
      },
      {
        path: 'discover',
        element: <Discover />,
      },
      {
        path: 'voice-plan',
        element: <VoicePlan />,
      },
      {
        path: 'guide',
        element: <Guide />,
      },
      {
        path: 'profile',
        element: <Profile />,
      },
    ],
  },
])

export default router
