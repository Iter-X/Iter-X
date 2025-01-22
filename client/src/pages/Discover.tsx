import { useState, useEffect } from 'react'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Card, CardContent } from '@/components/ui/card'
import { Star } from 'lucide-react'
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from '@/components/ui/carousel'

interface Place {
  id: string
  name: string
  type: string
  rating: number
  distance: string
  imageUrl: string
  address: string
  tags: string[]
}

const mockPlaces: Place[] = [
  {
    id: '1',
    name: '西湖',
    type: '景点',
    rating: 4.8,
    distance: '2.5km',
    imageUrl: 'https://picsum.photos/300/200',
    address: '浙江省杭州市西湖区',
    tags: ['自然风光', '文化古迹', '网红打卡'],
  },
  {
    id: '2',
    name: '南宋御街',
    type: '历史街区',
    rating: 4.5,
    distance: '1.2km',
    imageUrl: 'https://picsum.photos/300/200',
    address: '浙江省杭州市上城区',
    tags: ['历史遗迹', '美食', '购物'],
  },
  {
    id: '3',
    name: '灵隐寺',
    type: '寺庙',
    rating: 4.7,
    distance: '3.8km',
    imageUrl: 'https://picsum.photos/300/200',
    address: '浙江省杭州市西湖区灵隐路1号',
    tags: ['佛教文化', '古迹', '祈福'],
  },
  {
    id: '4',
    name: '河坊街',
    type: '商业街',
    rating: 4.4,
    distance: '1.5km',
    imageUrl: 'https://picsum.photos/300/200',
    address: '浙江省杭州市上城区河坊街',
    tags: ['美食', '购物', '文化'],
  },
]

const Discover = () => {
  const [loading, setLoading] = useState(true)
  const [places, setPlaces] = useState<Place[]>([])
  const [searchQuery, setSearchQuery] = useState('')
  const [activeTab, setActiveTab] = useState<'最新' | '热度' | '精选'>('精选')

  useEffect(() => {
    // 模拟API请求
    const fetchNearbyPlaces = async () => {
      try {
        await new Promise((resolve) => setTimeout(resolve, 1000))
        setPlaces(mockPlaces)
      } catch (error) {
        console.error('获取周边地点失败:', error)
      } finally {
        setLoading(false)
      }
    }

    fetchNearbyPlaces()
  }, [])

  const filteredPlaces = places.filter(
    (place) =>
      place.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
      place.type.toLowerCase().includes(searchQuery.toLowerCase()) ||
      place.tags.some((tag) =>
        tag.toLowerCase().includes(searchQuery.toLowerCase())
      )
  )

  return (
    <div className='h-screen bg-gray-50 overflow-auto'>
      <div className='max-w-4xl mx-auto'>
        {/* 搜索区域 */}
        <div className='sticky top-0 bg-white z-10 p-0 shadow-sm'>
          <div className='flex gap-2 mb-4'>
            <Input
              placeholder='搜索地点、类型或标签'
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              className='flex-1'
            />
            <Button variant='outline'>筛选</Button>
          </div>
          <div className='flex gap-4'>
            <button
              className={`text-sm ${
                activeTab === '最新'
                  ? 'text-blue-500 font-medium'
                  : 'text-gray-500'
              }`}
              onClick={() => setActiveTab('最新')}
            >
              最新
            </button>
            <button
              className={`text-sm ${
                activeTab === '热度'
                  ? 'text-blue-500 font-medium'
                  : 'text-gray-500'
              }`}
              onClick={() => setActiveTab('热度')}
            >
              热度
            </button>
          </div>
        </div>

        {/* 内容区域 */}
        <div className='p-0'>
          {loading ? (
            <div className='text-center py-8'>加载中...</div>
          ) : (
            <div className='space-y-4'>
              {/* 精选内容 */}
              <div className='space-y-4'>
                <div className='flex gap-4 border-b'>
                  {['精选', '最新', '热度'].map((tab) => (
                    <button
                      key={tab}
                      className={`pb-2 px-1 text-sm ${
                        activeTab === tab
                          ? 'text-blue-500 font-medium border-b-2 border-blue-500'
                          : 'text-gray-500'
                      }`}
                      onClick={() =>
                        setActiveTab(tab as '精选' | '最新' | '热度')
                      }
                    >
                      {tab}
                    </button>
                  ))}
                </div>
                <Carousel className='w-full'>
                  <CarouselContent>
                    {filteredPlaces.slice(0, 4).map((place) => (
                      <CarouselItem key={place.id}>
                        <div className='bg-white rounded-lg overflow-hidden shadow-sm'>
                          <img
                            src={place.imageUrl}
                            alt={place.name}
                            className='w-full h-40 object-cover'
                          />
                          <div className='p-4'>
                            <div className='inline-block px-2 py-1 bg-blue-50 text-blue-600 text-xs rounded-full mb-2'>
                              精选
                            </div>
                            <h3 className='text-lg font-semibold mb-1'>
                              {place.name}
                            </h3>
                            <p className='text-sm text-gray-500'>
                              {place.address}
                            </p>
                          </div>
                        </div>
                      </CarouselItem>
                    ))}
                  </CarouselContent>
                  <CarouselPrevious className='left-2' />
                  <CarouselNext className='right-2' />
                </Carousel>
              </div>

              {/* 瀑布流列表 */}
              <div className='columns-1 md:columns-2 gap-4 space-y-4'>
                {filteredPlaces.map((place) => (
                  <Card
                    key={place.id}
                    className='overflow-hidden hover:shadow-md transition-shadow'
                  >
                    <img
                      src={place.imageUrl}
                      alt={place.name}
                      className='w-full h-40 object-cover'
                    />
                    <CardContent className='p-4'>
                      <div className='flex justify-between items-start mb-2'>
                        <div>
                          <h3 className='text-base font-semibold'>
                            {place.name}
                          </h3>
                          <p className='text-sm text-gray-500'>{place.type}</p>
                        </div>
                        <div className='flex items-center'>
                          <Star className='w-4 h-4 text-yellow-400 fill-current' />
                          <span className='ml-1 text-sm'>{place.rating}</span>
                        </div>
                      </div>
                      <p className='text-sm text-gray-600 mb-2'>
                        {place.address}
                      </p>
                      <div className='flex items-center justify-between'>
                        <div className='flex gap-2'>
                          {place.tags.map((tag) => (
                            <span
                              key={tag}
                              className='px-2 py-1 bg-gray-50 text-xs rounded-full text-gray-600'
                            >
                              {tag}
                            </span>
                          ))}
                        </div>
                        <span className='text-xs text-gray-500'>
                          {place.distance}
                        </span>
                      </div>
                    </CardContent>
                  </Card>
                ))}
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  )
}

export default Discover
