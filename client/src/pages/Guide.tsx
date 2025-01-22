import { useState, useEffect } from 'react'
import { Card, CardContent } from '@/components/ui/card'
import { Star } from 'lucide-react'

interface GuideItem {
  id: string
  title: string
  author: string
  rating: number
  views: number
  tags: string[]
  coverImage: string
  summary: string
}

const mockGuides: GuideItem[] = [
  {
    id: '1',
    title: '杭州西湖三日游完全攻略',
    author: '旅行达人',
    rating: 4.9,
    views: 12345,
    tags: ['西湖', '人文', '美食'],
    coverImage: 'https://picsum.photos/300/200',
    summary: '详细介绍西湖周边景点、美食推荐和交通建议',
  },
  {
    id: '2',
    title: '灵隐寺一日游攻略',
    author: '文化探索者',
    rating: 4.7,
    views: 8765,
    tags: ['寺庙', '文化', '素食'],
    coverImage: 'https://picsum.photos/300/200',
    summary: '深度探访灵隐寺，感受佛教文化魅力',
  },
]

const Guide = () => {
  const [loading, setLoading] = useState(true)
  const [guides, setGuides] = useState<GuideItem[]>([])

  useEffect(() => {
    const fetchGuides = async () => {
      try {
        await new Promise((resolve) => setTimeout(resolve, 1000))
        setGuides(mockGuides)
      } catch (error) {
        console.error('获取攻略列表失败:', error)
      } finally {
        setLoading(false)
      }
    }

    fetchGuides()
  }, [])

  return (
    <div className='p-4 bg-gray-50 min-h-screen'>
      <div className='max-w-4xl mx-auto'>
        <h1 className='text-2xl font-bold mb-6'>精选攻略</h1>
        {loading ? (
          <div className='text-center py-8'>加载中...</div>
        ) : (
          <div className='grid gap-6'>
            {guides.map((guide) => (
              <Card
                key={guide.id}
                className='overflow-hidden hover:shadow-lg transition-shadow'
              >
                <div className='flex flex-col md:flex-row'>
                  <img
                    src={guide.coverImage}
                    alt={guide.title}
                    className='w-full md:w-48 h-48 object-cover'
                  />
                  <CardContent className='flex-1 p-4'>
                    <div className='flex justify-between items-start mb-2'>
                      <div>
                        <h3 className='text-xl font-semibold'>{guide.title}</h3>
                        <p className='text-sm text-gray-500'>
                          作者: {guide.author}
                        </p>
                      </div>
                      <div className='flex items-center gap-4'>
                        <div className='flex items-center'>
                          <Star className='w-4 h-4 text-yellow-400 fill-current' />
                          <span className='ml-1 text-sm'>{guide.rating}</span>
                        </div>
                        <span className='text-sm text-gray-500'>
                          {guide.views} 浏览
                        </span>
                      </div>
                    </div>
                    <p className='text-gray-600 mb-4'>{guide.summary}</p>
                    <div className='flex gap-2'>
                      {guide.tags.map((tag) => (
                        <span
                          key={tag}
                          className='px-2 py-1 bg-gray-100 text-xs rounded-full'
                        >
                          {tag}
                        </span>
                      ))}
                    </div>
                  </CardContent>
                </div>
              </Card>
            ))}
          </div>
        )}
      </div>
    </div>
  )
}

export default Guide
