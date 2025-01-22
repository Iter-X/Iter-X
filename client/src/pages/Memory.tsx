import { useState, useEffect } from 'react'
import { Button } from '@/components/ui/button'
import { Card, CardContent } from '@/components/ui/card'
import { PlusCircle } from 'lucide-react'

interface Memory {
  id: string
  title: string
  date: string
  location: string
  imageUrl: string
  description: string
}

const mockMemories: Memory[] = [
  {
    id: '1',
    title: '西湖游船',
    date: '2024-01-15',
    location: '杭州西湖',
    imageUrl: 'https://picsum.photos/300/200',
    description: '和家人一起游西湖，看断桥残雪，品龙井茶',
  },
  {
    id: '2',
    title: '灵隐寺祈福',
    date: '2024-01-16',
    location: '灵隐寺',
    imageUrl: 'https://picsum.photos/300/200',
    description: '参观灵隐寺，感受佛教文化的魅力',
  },
]

const Memory = () => {
  const [loading, setLoading] = useState(true)
  const [memories, setMemories] = useState<Memory[]>([])

  useEffect(() => {
    const fetchMemories = async () => {
      try {
        await new Promise((resolve) => setTimeout(resolve, 1000))
        setMemories(mockMemories)
      } catch (error) {
        console.error('获取回忆列表失败:', error)
      } finally {
        setLoading(false)
      }
    }

    fetchMemories()
  }, [])

  return (
    <div className='p-4 bg-gray-50 min-h-screen'>
      <div className='max-w-4xl mx-auto'>
        <div className='flex justify-between items-center mb-6'>
          <h1 className='text-2xl font-bold'>我的回忆</h1>
          <Button className='flex items-center gap-2'>
            <PlusCircle className='w-4 h-4' />
            <span>添加回忆</span>
          </Button>
        </div>

        {loading ? (
          <div className='text-center py-8'>加载中...</div>
        ) : (
          <div className='grid grid-cols-1 md:grid-cols-2 gap-6'>
            {memories.map((memory) => (
              <Card
                key={memory.id}
                className='overflow-hidden hover:shadow-lg transition-shadow'
              >
                <img
                  src={memory.imageUrl}
                  alt={memory.title}
                  className='w-full h-48 object-cover'
                />
                <CardContent className='p-4'>
                  <h3 className='text-xl font-semibold mb-2'>{memory.title}</h3>
                  <div className='text-sm text-gray-500 space-y-1 mb-3'>
                    <p>{memory.date}</p>
                    <p>{memory.location}</p>
                  </div>
                  <p className='text-gray-600'>{memory.description}</p>
                </CardContent>
              </Card>
            ))}
          </div>
        )}
      </div>
    </div>
  )
}

export default Memory
