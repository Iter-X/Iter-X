import { useState } from 'react'
import { Button } from '@/components/ui/button'
import { Card, CardContent } from '@/components/ui/card'
import { Mic, MicOff } from 'lucide-react'

interface PlanItem {
  id: string
  time: string
  activity: string
  location: string
  description: string
}

const mockPlan: PlanItem[] = [
  {
    id: '1',
    time: '09:00',
    activity: '游览西湖',
    location: '杭州西湖',
    description: '乘船游览西湖，欣赏湖光山色',
  },
  {
    id: '2',
    time: '11:30',
    activity: '品尝龙井茶',
    location: '龙井村',
    description: '参观茶园，品尝正宗龙井茶',
  },
]

const VoicePlan = () => {
  const [recording, setRecording] = useState(false)
  const [loading, setLoading] = useState(false)
  const [plan, setPlan] = useState<PlanItem[]>([])

  const startRecording = async () => {
    setRecording(true)
    // TODO: 实现语音录制逻辑
  }

  const stopRecording = async () => {
    setRecording(false)
    setLoading(true)
    try {
      // 模拟API调用
      await new Promise((resolve) => setTimeout(resolve, 1500))
      setPlan(mockPlan)
    } catch (error) {
      console.error('生成行程规划失败:', error)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className='p-4 h-screen bg-gray-50'>
      <div className='max-w-4xl mx-auto space-y-6'>
        <div className='text-center space-y-4'>
          <h1 className='text-2xl font-bold'>语音规划</h1>
          <p className='text-gray-500'>
            请说出您想要的行程安排，我们将为您生成个性化的旅行计划
          </p>
          <Button
            size='lg'
            className='w-32 h-32 rounded-full'
            variant={recording ? 'destructive' : 'default'}
            onClick={recording ? stopRecording : startRecording}
          >
            {recording ? (
              <MicOff className='w-12 h-12' />
            ) : (
              <Mic className='w-12 h-12' />
            )}
          </Button>
          <p className='text-sm text-gray-500'>
            {recording ? '点击停止录音' : '点击开始录音'}
          </p>
        </div>

        {loading ? (
          <div className='text-center py-8'>生成行程规划中...</div>
        ) : plan.length > 0 ? (
          <div className='space-y-4'>
            <h2 className='text-xl font-semibold'>为您规划的行程</h2>
            <div className='grid gap-4'>
              {plan.map((item) => (
                <Card key={item.id}>
                  <CardContent className='p-4'>
                    <div className='flex items-start gap-4'>
                      <div className='text-lg font-medium text-gray-500'>
                        {item.time}
                      </div>
                      <div className='flex-1'>
                        <h3 className='text-lg font-semibold mb-1'>
                          {item.activity}
                        </h3>
                        <p className='text-sm text-gray-500 mb-2'>
                          {item.location}
                        </p>
                        <p className='text-gray-600'>{item.description}</p>
                      </div>
                    </div>
                  </CardContent>
                </Card>
              ))}
            </div>
          </div>
        ) : null}
      </div>
    </div>
  )
}

export default VoicePlan
