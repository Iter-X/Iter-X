import { Outlet, useLocation, useNavigate } from 'react-router-dom'
import { Home, Compass, Mic, BookOpen, User } from 'lucide-react'

const Layout = () => {
  const navigate = useNavigate()
  const location = useLocation()

  const tabs = [
    { key: 'memory', label: '回忆', icon: <Home className='w-6 h-6' /> },
    { key: 'discover', label: '发现', icon: <Compass className='w-6 h-6' /> },
    {
      key: 'voice-plan',
      label: '',
      icon: <Mic className='w-6 h-6' />,
    },
    { key: 'guide', label: '攻略', icon: <BookOpen className='w-6 h-6' /> },
    { key: 'profile', label: '我的', icon: <User className='w-6 h-6' /> },
  ]

  return (
    <div className='flex flex-col h-screen bg-gray-50'>
      <main className='flex-1 overflow-y-auto pt-16 pb-20 px-4'>
        <Outlet />
      </main>
      <nav className='fixed bottom-0 left-0 right-0 flex items-center justify-around py-4 bg-white border-t border-gray-100 shadow-lg'>
        {tabs.map((tab) => {
          const isActive = location.pathname.includes(tab.key)
          const isVoicePlan = tab.key === 'voice-plan'
          return (
            <div
              key={tab.key}
              className={`flex flex-col items-center ${
                isVoicePlan ? '-mt-8' : ''
              } cursor-pointer transition-colors duration-200 ease-in-out`}
              onClick={() => navigate(tab.key)}
            >
              {isVoicePlan ? (
                <div className='w-16 h-16 rounded-full bg-blue-500 flex items-center justify-center text-white shadow-xl transform -translate-y-6 hover:bg-blue-600 transition-all duration-200 ease-in-out'>
                  <span className='text-2xl'>{tab.icon}</span>
                </div>
              ) : (
                <>
                  <span
                    className={`text-2xl mb-1 ${
                      isActive ? 'text-blue-500' : 'text-gray-400'
                    } transition-colors duration-200 ease-in-out`}
                  >
                    {tab.icon}
                  </span>
                  <span
                    className={`text-xs ${
                      isActive ? 'text-blue-500 font-medium' : 'text-gray-500'
                    } transition-colors duration-200 ease-in-out`}
                  >
                    {tab.label}
                  </span>
                </>
              )}
            </div>
          )
        })}
      </nav>
    </div>
  )
}

export default Layout
