import { useState } from 'react'
import { Button } from '@/components/ui/button'
import { Card, CardContent } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Switch } from '@/components/ui/switch'
import { User, Settings, LogOut, Bell, Moon, Sun, Globe } from 'lucide-react'

interface UserProfile {
  avatar: string
  username: string
  email: string
  joinDate: string
  memoryCount: number
  guideCount: number
  notifications: boolean
  darkMode: boolean
  language: string
}

const mockProfile: UserProfile = {
  avatar: 'https://picsum.photos/200',
  username: '旅行者小明',
  email: 'xiaoming@example.com',
  joinDate: '2024-01-01',
  memoryCount: 12,
  guideCount: 5,
  notifications: true,
  darkMode: false,
  language: '简体中文',
}

const Profile = () => {
  const [profile, setProfile] = useState<UserProfile>(mockProfile)
  const [editing, setEditing] = useState(false)

  const handleLogout = () => {
    localStorage.removeItem('isLoggedIn')
    window.location.href = '/login'
  }

  return (
    <div className='p-0 m-0 bg-gray-50 min-h-screen'>
      <div className='mx-auto max-w-4xl space-y-6'>
        <Card>
          <CardContent className='p-6'>
            <div className='flex items-center gap-4 mb-6'>
              <img
                src={profile.avatar}
                alt={profile.username}
                className='rounded-full object-cover h-12 w-12'
              />
              <div className='flex-1'>
                <h2 className='text-2xl font-bold'>{profile.username}</h2>
                <p className='text-gray-500'>加入时间：{profile.joinDate}</p>
              </div>
            </div>

            <div className='grid grid-cols-2 gap-4 mb-6'>
              <div className='text-center p-4 bg-gray-100 rounded-lg'>
                <p className='text-2xl font-bold'>{profile.memoryCount}</p>
                <p className='text-gray-500'>回忆</p>
              </div>
              <div className='text-center p-4 bg-gray-100 rounded-lg'>
                <p className='text-2xl font-bold'>{profile.guideCount}</p>
                <p className='text-gray-500'>攻略</p>
              </div>
            </div>

            {editing ? (
              <div className='space-y-4'>
                <div className='space-y-2'>
                  <Label htmlFor='username'>用户名</Label>
                  <Input
                    id='username'
                    value={profile.username}
                    onChange={(e) =>
                      setProfile({ ...profile, username: e.target.value })
                    }
                  />
                </div>
                <div className='space-y-2'>
                  <Label htmlFor='email'>邮箱</Label>
                  <Input
                    id='email'
                    type='email'
                    value={profile.email}
                    onChange={(e) =>
                      setProfile({ ...profile, email: e.target.value })
                    }
                  />
                </div>
                <div className='flex justify-end gap-2'>
                  <Button variant='outline' onClick={() => setEditing(false)}>
                    取消
                  </Button>
                  <Button onClick={() => setEditing(false)}>保存</Button>
                </div>
              </div>
            ) : (
              <div className='space-y-6'>
                <div className='space-y-4'>
                  <h3 className='text-lg font-semibold'>设置</h3>
                  <div className='space-y-4'>
                    <div className='flex items-center justify-between'>
                      <div className='flex items-center gap-2'>
                        <Bell className='w-4 h-4' />
                        <span>通知提醒</span>
                      </div>
                      <Switch
                        checked={profile.notifications}
                        onCheckedChange={(checked) =>
                          setProfile({ ...profile, notifications: checked })
                        }
                      />
                    </div>
                    <div className='flex items-center justify-between'>
                      <div className='flex items-center gap-2'>
                        {profile.darkMode ? (
                          <Moon className='w-4 h-4' />
                        ) : (
                          <Sun className='w-4 h-4' />
                        )}
                        <span>深色模式</span>
                      </div>
                      <Switch
                        checked={profile.darkMode}
                        onCheckedChange={(checked) =>
                          setProfile({ ...profile, darkMode: checked })
                        }
                      />
                    </div>
                    <div className='flex items-center justify-between'>
                      <div className='flex items-center gap-2'>
                        <Globe className='w-4 h-4' />
                        <span>语言</span>
                      </div>
                      <select
                        value={profile.language}
                        onChange={(e) =>
                          setProfile({ ...profile, language: e.target.value })
                        }
                        className='border rounded-md px-2 py-1'
                      >
                        <option value='简体中文'>简体中文</option>
                        <option value='English'>English</option>
                      </select>
                    </div>
                  </div>
                </div>
                <div className='space-y-4'>
                  <Button
                    variant='outline'
                    className='w-full flex items-center gap-2'
                    onClick={() => setEditing(true)}
                  >
                    <Settings className='w-4 h-4' />
                    编辑资料
                  </Button>
                  <Button
                    variant='destructive'
                    className='w-full flex items-center gap-2'
                    onClick={handleLogout}
                  >
                    <LogOut className='w-4 h-4' />
                    退出登录
                  </Button>
                </div>
              </div>
            )}
          </CardContent>
        </Card>
      </div>
    </div>
  )
}

export default Profile
