import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { Button } from '@/components/ui/button'
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
  CardDescription,
} from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Toast } from '@/components/ui/toast'

const Login = () => {
  const navigate = useNavigate()
  const [loading, setLoading] = useState(false)

  const onFinish = async (values: any) => {
    setLoading(true)
    try {
      // TODO: 实现实际的登录逻辑
      console.log('登录信息:', values)
      navigate('/memory')
    } catch (error) {
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className='min-h-screen flex items-center justify-center bg-gray-100 px-4'>
      <Card className='w-full max-w-md'>
        <CardHeader>
          <CardTitle className='text-center text-3xl'>旅行攻略</CardTitle>
          <CardDescription className='text-center'>
            开启您的智能旅行体验
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form
            onSubmit={(e) => {
              e.preventDefault()
              const formData = new FormData(e.currentTarget)
              onFinish({
                username: formData.get('username'),
                password: formData.get('password'),
              })
            }}
          >
            <div className='space-y-4'>
              <div className='space-y-2'>
                <Label htmlFor='username'>用户名</Label>
                <Input id='username' name='username' type='text' required />
              </div>
              <div className='space-y-2'>
                <Label htmlFor='password'>密码</Label>
                <Input id='password' name='password' type='password' required />
              </div>
              <Button type='submit' className='w-full' disabled={loading}>
                {loading ? '登录中...' : '登录'}
              </Button>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  )
}

export default Login
