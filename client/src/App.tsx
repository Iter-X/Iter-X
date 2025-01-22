import { useEffect } from 'react'
import { RouterProvider } from 'react-router-dom'
import router from './router'

function App() {
  useEffect(() => {
    // 检查用户是否已登录，如果未登录则重定向到登录页面
    const isLoggedIn = localStorage.getItem('isLoggedIn')
    if (!isLoggedIn && window.location.pathname !== '/login') {
      window.location.href = '/login'
    }
  }, [])

  return <RouterProvider router={router} />
}
export default App
