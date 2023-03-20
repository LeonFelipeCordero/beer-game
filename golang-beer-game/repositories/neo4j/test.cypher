CREATE (b:BoardNode {
  id:         '1234',
  UUID:       '1234',
  name:       'test',
  state:      'PENDING',
  full:       false,
  finished:   false,
  created_at: localDateTime()
})
RETURN b;

MATCH (b:BoardNode)
  WHERE b.UUID = '1234'
CREATE (p:PlayerNode {
  id:           '1234',
  UUID:         '1234',
  name:         'RETAILER',
  role:         'RETAILER',
  stock:        40,
  weekly_order: 40,
  last_order:   40,
  cpu:          false
})-[r:plays_in]->(b)
RETURN p;

MATCH (b:BoardNode)
  WHERE b.UUID = '1234'
CREATE (p:PlayerNode {
  id:           '1235',
  UUID:         '1235',
  name:         'WHOLESALER',
  role:         'WHOLESALER',
  stock:        240,
  weekly_order: 240,
  last_order:   240,
  cpu:          false
})-[r:plays_in]->(b)
RETURN p;

MATCH (b:BoardNode)
  WHERE b.UUID = '1234'
CREATE (p:PlayerNode {
  id:           '1236',
  UUID:         '1236',
  name:         'FACTORY',
  role:         'FACTORY',
  stock:        2400,
  weekly_order: 2400,
  last_order:   2400,
  cpu:          false
})-[r:plays_in]->(b)
RETURN p;

MATCH (r:PlayerNode {UUID: '1234'})
MATCH (s:PlayerNode {UUID: '1235'})
WITH r, s
CREATE (o:OrderNode {
  id:              '1234',
  UUID:            '1234',
  amount:          40,
  original_amount: 40,
  state:           'PENDING',
  order_type:      'PLAYER',
  created_at:      localDateTime()
})
CREATE(o)<-[or:ordered]-(r),
      (o)-[re:received]->(s)
RETURN o


MATCH (n)
OPTIONAL MATCH (n)-[r:plays_in]->(m)
DELETE n, r, m