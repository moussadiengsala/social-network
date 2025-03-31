package internals

const (
	QUERY_GETTING_All_POST = `SELECT
    p.id,
    p.content,
    p.author_id,
    p.image,
    p.privacy,
    COUNT(DISTINCT c.id) AS comment_count,
    u.first_name,
    u.last_name,
    u.username AS author_username,
    u.avatar AS author_avatar,

    COUNT(DISTINCT l.id) AS like_count,
    COUNT(DISTINCT CASE WHEN l.author_id = %d THEN l.id END) > 0 AS like_status,
    p.created_at

	FROM
		post p
	LEFT JOIN
		post_visibility pv ON p.id = pv.post_id AND pv.user_id = %d
	LEFT JOIN
		post_like l ON p.id = l.entries_id
	LEFT JOIN
		comments c ON p.id = c.entrie_id
	LEFT JOIN
		user u ON p.author_id = u.id
	LEFT JOIN
		follower f ON (f.follower_id = p.author_id AND f.following_id = '%d' AND f.status = 'accept') OR (f.follower_id = '%d' AND f.following_id =  p.author_id AND f.status = 'accept')

	WHERE
		(
			p.privacy = 'public'
			OR (p.privacy = 'private' AND (p.author_id = %d OR f.follower_id IS NOT NULL OR f.following_id IS NOT NULL))
			OR (p.privacy = 'almost_private' AND p.author_id = %d)
			OR (p.privacy = 'almost_private' AND pv.user_id = %d)
		)
	GROUP BY
		p.id
	ORDER BY
		p.created_at DESC
	LIMIT %s OFFSET %s;
	`

	QUERY_GETTING_LIKED_POST = `SELECT
		p.id,
		p.content,
		p.author_id,
		p.image,
		p.privacy,
		COUNT(DISTINCT c.id) AS comment_count,
		u.first_name,
		u.last_name,
		u.username AS author_username,
		u.avatar AS author_avatar,
		COUNT(DISTINCT l.id) AS like_count,
		COUNT(DISTINCT CASE WHEN l.author_id = %d THEN l.id END) > 0 AS like_status,
		p.created_at

		FROM
			post p
		LEFT JOIN
			post_visibility pv ON p.id = pv.post_id AND pv.user_id = %d
		LEFT JOIN
			post_like l ON p.id = l.entries_id
		LEFT JOIN
			comments c ON c.entrie_id = p.id
		LEFT JOIN
			user u ON p.author_id = u.id
		LEFT JOIN
			follower f ON (f.follower_id = p.author_id AND f.following_id = '%d' AND f.status = 'accept') OR (f.follower_id = '%d' AND f.following_id =  p.author_id AND f.status = 'accept')

		WHERE
			l.author_id = %d AND (
				p.privacy = 'public'
				OR (p.privacy = 'private' AND (p.author_id = %d OR f.follower_id IS NOT NULL OR f.following_id IS NOT NULL))
				OR (p.privacy = 'almost_private' AND p.author_id = %d)
				OR (p.privacy = 'almost_private' AND pv.user_id = %d)
			)
		GROUP BY
			p.id
		ORDER BY
			p.created_at DESC
		LIMIT %s OFFSET %s;
	`

	QUERYGETTINGALLNOTIFICATIONS = `
	SELECT
    n.id,
    n.receiver_id,
    n.sender_id,
    u.username AS sender_username,
    n.notification_type,
	n.group_id,
    n.status,
    n.created_at
	FROM
		notifications n
	JOIN
		user u ON n.sender_id = u.id
	WHERE
		n.receiver_id = %d
	ORDER BY
		n.created_at DESC;`

	QUERY_GETTING_COMMENTED_POST = `SELECT
		p.id,
		p.content,
		p.author_id,
		p.image,
		p.privacy,
		COUNT(DISTINCT c.id) AS comment_count,
		u.first_name,
		u.last_name,
		u.username AS author_username,
		u.avatar AS author_avatar,
		COUNT(DISTINCT l.id) AS like_count,
		COUNT(DISTINCT CASE WHEN l.author_id = %d THEN l.id END) > 0 AS like_status,
		p.created_at

		FROM
			post p
		LEFT JOIN
			post_visibility pv ON p.id = pv.post_id AND pv.user_id = %d
		LEFT JOIN
			post_like l ON p.id = l.entries_id
		LEFT JOIN
			comments c ON c.entrie_id = p.id
		LEFT JOIN
			user u ON p.author_id = u.id
		LEFT JOIN
			follower f ON (f.follower_id = p.author_id AND f.following_id = '%d' AND f.status = 'accept') OR (f.follower_id = '%d' AND f.following_id =  p.author_id AND f.status = 'accept')

		WHERE
			c.author_id = %d AND (
				p.privacy = 'public'
				OR (p.privacy = 'private' AND (p.author_id = %d OR f.follower_id IS NOT NULL OR f.following_id IS NOT NULL))
				OR (p.privacy = 'almost_private' AND p.author_id = %d)
				OR (p.privacy = 'almost_private' AND pv.user_id = %d)
			)
		GROUP BY
			p.id
		ORDER BY
			p.created_at DESC
		LIMIT %s OFFSET %s;
	`

	QUERY_GETTING_OWN_POST = `SELECT
		p.id,
		p.content,
		p.author_id,
		p.image,
		p.privacy,
		COUNT(DISTINCT c.id) AS comment_count,
		u.first_name,
		u.last_name,
		u.username AS author_username,
		u.avatar AS author_avatar,
		COUNT(DISTINCT l.id) AS like_count,
		COUNT(DISTINCT CASE WHEN l.author_id = %d THEN l.id END) > 0 AS like_status,
		p.created_at

		FROM
			post p
		LEFT JOIN
			post_visibility pv ON p.id = pv.post_id AND pv.user_id = %d
		LEFT JOIN
			post_like l ON p.id = l.entries_id
		LEFT JOIN
			comments c ON p.id = c.entrie_id
		LEFT JOIN
			user u ON p.author_id = u.id
		LEFT JOIN
			follower f ON (f.follower_id = p.author_id AND f.following_id = '%d' AND f.status = 'accept') OR (f.follower_id = '%d' AND f.following_id =  p.author_id AND f.status = 'accept')

		WHERE
			p.author_id = %d AND (
				p.privacy = 'public'
				OR (p.privacy = 'private' AND (p.author_id = %d OR f.follower_id IS NOT NULL OR f.following_id IS NOT NULL))
				OR (p.privacy = 'almost_private' AND pv.user_id = %d)
			)
		GROUP BY
			p.id
		ORDER BY
			p.created_at DESC
		LIMIT %s OFFSET %s;
	`

	QUERY_GETTING_SINGLE_POST = `
	SELECT
    p.id,
    p.content,
    p.author_id,
    p.image,
    p.privacy,
    COUNT(DISTINCT c.id) AS comment_count,
    u.first_name,
    u.last_name,
    u.username AS author_username,
    u.avatar AS author_avatar,
    COUNT(DISTINCT l.id) AS like_count,
    COUNT(DISTINCT CASE WHEN l.author_id = %d THEN l.id END) > 0 AS like_status,
    p.created_at
	FROM
		post p
	LEFT JOIN
		post_visibility pv ON p.id = pv.post_id AND pv.user_id = %d
	LEFT JOIN
		post_like l ON p.id = l.entries_id
	LEFT JOIN
		comments c ON p.id = c.entrie_id
	LEFT JOIN
		user u ON p.author_id = u.id
	LEFT JOIN
		follower f ON (f.follower_id = p.author_id AND f.following_id = '%d' AND f.status = 'accept') OR (f.follower_id = '%d' AND f.following_id =  p.author_id AND f.status = 'accept')

	WHERE
    p.id = %d
    AND (
        p.privacy = 'public'
        OR (p.privacy = 'private' AND (p.author_id = %d OR f.follower_id IS NOT NULL OR f.following_id IS NOT NULL))
        OR (p.privacy = 'almost_private' AND p.author_id = %d)
        OR (p.privacy = 'almost_private' AND pv.user_id = %d)
    )
	GROUP BY
    p.id;

	`

	QUERY_GETTING_All_COMMENTS = `SELECT
	c.id,
	c.entrie_id,
	c.author_id,
	c.content,
	c.image,
	c.created_at,
	u.avatar,
	u.username,
	COUNT(DISTINCT cl.id) AS like_count,

	COUNT(DISTINCT CASE WHEN cl.author_id = %d THEN cl.id END) > 0 AS like_status

	FROM
		comments c
	LEFT JOIN
		comment_like cl ON c.id = cl.entries_id
	LEFT JOIN
		user u ON c.author_id = u.id
	WHERE c.entrie_id = %d
	GROUP BY c.entrie_id, c.author_id, c.content, c.image, c.created_at, u.avatar, u.username
	ORDER BY
		c.created_at DESC -- Order by created_at in descending order (most recent first)
	LIMIT %s OFFSET %s`

	QUERYGETTINGSINGLECOMMENTS = `SELECT
    c.id,
    c.entrie_id,
    c.author_id,
    c.content,
    c.image,
    c.created_at,
    u.avatar AS author_avatar,
    u.username AS author_username,
    COUNT(DISTINCT cl.id) AS like_count,
    COUNT(DISTINCT CASE WHEN cl.author_id = ? THEN cl.id END) > 0 AS like_status

	FROM
		comments c
	LEFT JOIN
		comment_like cl ON c.id = cl.comment_id
	LEFT JOIN
		user u ON c.author_id = u.id
	WHERE
		c.entrie_id = ?
	GROUP BY c.id, c.entrie_id, c.author_id, c.content, c.image, c.created_at, u.avatar, u.username
	ORDER BY
		c.created_at DESC;

	`
	QUERY_GETTING_USER = `SELECT
		%s, u.created_at,

		'' AS following_status,
		COUNT(DISTINCT f.id) AS followers,
		COUNT(DISTINCT fe.id) AS following,
		0 AS likes_count,
		0 AS post_count,
		0 AS online,
		0 AS unread_message,
		NULL AS last_activity

		FROM
			user u
		LEFT JOIN
			follower f ON u.id = f.follower_id
		LEFT JOIN
			follower fe ON u.id = fe.following_id
		WHERE u.username = '%s' OR u.email = '%s'
		GROUP BY %s, u.created_at;
	`

	// Groups
	QUERY_GETTING_OTHER_GROUPS = `
		SELECT
			g.id,
			g.title,
			g.description,
			g.author_id,
			g.cover,
			g.created_at,

			u.username,
			u.first_name,
			u.last_name,
			u.avatar,

			COUNT(DISTINCT CASE WHEN gm.status = 'accepted' THEN gm.user_id END) AS member_count,
			COUNT(DISTINCT gp.id) AS post_count,
			COUNT(DISTINCT ge.id) AS event_count,
			0 AS unread_message,
			"" AS last_activity,
			0 AS is_part_of_group,
			(CASE
				WHEN EXISTS (
					SELECT gm.status
					FROM groupMembers gm
					WHERE gm.group_id = g.id AND gm.user_id = %d
				)
				THEN (SELECT gm.status FROM groupMembers gm WHERE gm.group_id = g.id AND gm.user_id = %d)
				ELSE 'null'
			END) AS status

		FROM
			groups g
		LEFT JOIN
			user u ON g.author_id = u.id
		LEFT JOIN
			groupMembers gm ON g.id = gm.group_id
		LEFT JOIN
			groupPosts gp ON g.id = gp.group_id
		LEFT JOIN
			groupEvents ge ON g.id = ge.group_id

		WHERE
			g.id NOT IN (SELECT gm.group_id FROM groupMembers gm WHERE gm.user_id = %d AND gm.status = 'accepted')

		GROUP BY
			g.id, g.title, g.description, g.author_id, g.cover, g.created_at, u.username, u.first_name, u.last_name, u.avatar
		LIMIT %s OFFSET %s;
	`

	QUERYGETTINGGROUPS = `SELECT
		g.id,
		g.title,
		g.description,
		g.author_id,
		g.cover,
		g.created_at,

		u.username,
		u.first_name,
		u.last_name,
		u.avatar,

		COUNT(DISTINCT CASE WHEN gm.status = 'accepted' THEN gm.user_id END) AS member_count,
		COUNT(DISTINCT gp.id) AS post_count,
		COUNT(DISTINCT ge.id) AS event_count,
		0 AS unread_message,
		COALESCE(MAX(COALESCE(gp.created_at, ge.created_at, gm.joined_at)), '') AS last_activity,
		(CASE
			WHEN EXISTS (
				SELECT 1
				FROM groupMembers gm
				WHERE gm.group_id = g.id AND gm.user_id = %d AND gm.status = 'accepted'
			) THEN 1 ELSE 0 END) AS is_part_of_group,
		(CASE
			WHEN EXISTS (
				SELECT gm.status
				FROM groupMembers gm
				WHERE gm.group_id = g.id AND gm.user_id = %d
			)
			THEN (SELECT gm.status FROM groupMembers gm WHERE gm.group_id = g.id AND gm.user_id = %d)
			ELSE 'null'
		END) AS status

		FROM
			groups g
		LEFT JOIN
			user u ON g.author_id = u.id
		LEFT JOIN
			groupMembers gm ON g.id = gm.group_id
		LEFT JOIN
			groupPosts gp ON g.id = gp.group_id
		LEFT JOIN
			groupEvents ge ON g.id = ge.group_id

		WHERE g.id = %d
		GROUP BY g.id, g.title, g.description, g.author_id, g.cover, g.created_at, u.username, u.first_name, u.last_name, u.avatar;
	`

	// Groups all Members
	QUERYGETTINGAllMEMBERSGROUPS = `SELECT
		u.id,
		u.first_name,
		u.last_name,
		u.email,
		u.username,
		u.avatar,
		m.role

		FROM
			groupMembers m
		LEFT JOIN
			user u ON m.user_id = u.id

		WHERE m.group_id = %d AND m.status = "accepted"
		GROUP BY u.id, u.first_name, u.last_name, u.email, u.username, u.avatar, m.role;
	`

	// Group Posts
	QUERY_GETTING_ALL_GROUP_POSTS = `SELECT
		gp.id,
		gp.group_id,
		gp.author_id,
		gp.content,
		gp.image,
		gp.created_at,

		u.first_name AS first_name,
		u.last_name AS last_name,
		u.username AS username,
		u.avatar AS avatar,

		(
			SELECT COUNT(*)
			FROM groupComments gc
			WHERE gc.post_id = gp.id
		) AS comment_count,

		g.title AS group_title,
		g.description AS group_description,
		g.cover AS group_cover,
		g.created_at AS group_created_at

		FROM groupPosts gp
		LEFT JOIN user u ON gp.author_id = u.id
		LEFT JOIN groups g ON gp.group_id = g.id
		LEFT JOIN groupComments gc ON gp.id = gc.post_id

		WHERE gp.group_id = %d
		ORDER BY gp.created_at DESC
		LIMIT %s OFFSET %s
	`

	QUERY_GETTING_GROUP_POSTS = `SELECT
		gp.id,
		gp.group_id,
		gp.author_id,
		gp.content,
		gp.image,
		gp.created_at,

		u.first_name AS 'first_name',
		u.last_name AS 'last_name',
		u.username AS 'username',
		u.avatar AS 'avatar',

		(
			SELECT COUNT(*)
			FROM groupComments gc
			WHERE gc.post_id = gp.id
		) AS comment_count,

		g.title AS 'group_title',
		g.description AS 'group_description',
		g.cover AS 'group_cover',
		g.created_at AS 'group_created_at'

		FROM groupPosts gp
		LEFT JOIN user u ON gp.author_id = u.id
		LEFT JOIN groups g ON gp.group_id = g.id
		LEFT JOIN groupComments gc ON gp.id = gc.post_id
		WHERE gp.id = ?;
	`

	// Group Comments
	QUERY_GETTING_ALL_GROUP_COMMENTS = `SELECT
		gc.id,
		gc.group_id,
		gc.author_id,
		gc.content,
		gc.image,
		gc.post_id,
		gc.created_at,

		u.first_name AS 'first_name',
		u.last_name AS 'last_name',
		u.username AS 'username',
		u.avatar AS 'avatar',

		g.title AS 'group_title',
		g.description AS 'group_description',
		g.cover AS 'group_cover',
		g.created_at AS 'group_created_at'

		FROM groupComments gc
		LEFT JOIN user u ON gc.author_id = u.id
		LEFT JOIN groups g ON gc.group_id = g.id
		LEFT JOIN groupPosts gp ON gp.id = gc.post_id

		WHERE gc.post_id = '%d'
		ORDER BY gc.created_at DESC
		LIMIT '%s' OFFSET '%s'
	`

	QUERY_GETTING_GROUP_COMMENTS = `SELECT
		gc.id,
		gc.group_id,
		gc.author_id,
		gc.content,
		gc.image,
		gc.post_id,
		gc.created_at,

		u.first_name AS 'first_name',
		u.last_name AS 'last_name',
		u.username AS 'username',
		u.avatar AS 'avatar',

		g.title AS 'group_title',
		g.description AS 'group_description',
		g.cover AS 'group_cover',
		g.created_at AS 'group_created_at'

		FROM groupComments gc
		LEFT JOIN user u ON gc.author_id = u.id
		LEFT JOIN groups g ON gc.group_id = g.id
		LEFT JOIN groupPosts gp ON gp.id = gc.post_id
		WHERE gc.id = '%d'
	`

	// Groups Events
	QUERY_GETTING_All_EVENTS_GROUPS = `SELECT
		e.id,
		e.title,
		e.description,
		u.id,
		e.datetime,
		u.first_name,
		u.last_name,
		u.email,
		u.username,
		u.avatar,


		(CASE
            WHEN EXISTS (
                SELECT response
                FROM groupEventResponses
                WHERE user_id = %d AND event_id = e.id
            )
            THEN (SELECT response FROM groupEventResponses WHERE user_id = %d AND event_id = e.id)
            ELSE 'null'
        END) AS response_status,
		COUNT(DISTINCT CASE WHEN r.response = 'going' THEN r.user_id END) AS going_count,
		COUNT(DISTINCT CASE WHEN r.response = 'not going' THEN r.user_id END) AS not_going_count

		FROM
			groupEvents e
		LEFT JOIN
			user u ON e.author_id = u.id
		LEFT JOIN
			groupEventResponses r ON r.event_id = e.id

		WHERE e.group_id = "%d"
		GROUP BY e.id, e.title, e.description, e.datetime, u.id, u.first_name, u.last_name, u.email, u.username, u.avatar
		LIMIT %s OFFSET %s;
	`

	QUERYGETTINGEVENTSGROUPS = `SELECT
		e.id,
		e.title,
		e.description,
		u.id,
		e.datetime,
		u.first_name,
		u.last_name,
		u.email,
		u.username,
		u.avatar,

		(CASE
            WHEN EXISTS (
                SELECT response
                FROM groupEventResponses
                WHERE user_id = %d AND event_id = e.id
            )
            THEN (SELECT response FROM groupEventResponses WHERE user_id = %d AND event_id = e.id)
            ELSE 'null'
        END) AS response_status,
		COUNT(DISTINCT CASE WHEN r.response = 'going' THEN r.user_id END) AS going_count,
		COUNT(DISTINCT CASE WHEN r.response = 'not going' THEN r.user_id END) AS not_going_count

		FROM
			groupEvents e
		LEFT JOIN
			user u ON e.author_id = u.id
		LEFT JOIN
			groupEventResponses r ON r.event_id = e.id

		WHERE e.group_id = "%d" AND e.id = "%d"
		GROUP BY e.id, e.title, e.description, e.datetime, u.id, u.first_name, u.last_name, u.email, u.username, u.avatar;
	`
	// event response
	QUERYGETTINGEVENTSGROUPSRESPONSE = `SELECT
		r.response,
		u.id,
		u.first_name,
		u.last_name,
		u.email,
		u.username,
		u.avatar

		FROM
			groupEventResponses r
		LEFT JOIN
			user u ON r.user_id = u.id

		WHERE r.group_id = "%d" AND r.event_id = "%d"
		GROUP BY r.response, u.id, u.first_name, u.last_name, u.email, u.username, u.avatar;
	`

	// get user profiles
	QUERY_GETTING_USER_PROFILE = `SELECT
		%s,
		u.created_at,

		(CASE
            WHEN EXISTS (
                SELECT f.status
                FROM follower f
                WHERE f.follower_id = u.id AND f.following_id = %d
            )
            THEN (SELECT f.status FROM follower f WHERE f.follower_id = u.id AND f.following_id = %d)
            ELSE 'not-followed'
        END) AS following_status,
		COUNT(DISTINCT f.id) AS followers,
		COUNT(DISTINCT fe.id) AS following,
		COUNT(DISTINCT pl.id) AS likes_count,
		COUNT(DISTINCT p.id) AS post_count,
		0 AS online,
		0 AS unread_message,
		NULL AS last_activity

		FROM
			user u
		LEFT JOIN
			follower f ON u.id = f.follower_id AND f.status = 'accept'
		LEFT JOIN
			follower fe ON u.id = fe.following_id AND fe.status = 'accept'
		LEFT JOIN
			post p ON p.author_id = u.id
		LEFT JOIN
			post_like pl ON pl.id = p.id

		WHERE u.id = %d
		GROUP BY %s, u.created_at;
	`

	// get all user
	QUERY_GETTING_All_USERS = `SELECT
		%s,
		u.created_at,

		(CASE
            WHEN EXISTS (
                SELECT f.status
                FROM follower
                WHERE follower_id = u.id AND following_id = %d
            )
            THEN (SELECT f.status FROM follower WHERE follower_id = u.id AND following_id = %d)
            ELSE 'not-followed'
        END) AS following_status,
		COUNT(DISTINCT f.id) AS followers,
		COUNT(DISTINCT fe.id) AS following,
		COUNT(DISTINCT pl.id) AS likes_count,
		COUNT(DISTINCT p.id) AS post_count,
		0 AS online,
		0 AS unread_message,
		null AS last_activity

		FROM
			user u
		LEFT JOIN
			follower f ON u.id = f.follower_id AND f.status = 'accept'
		LEFT JOIN
			follower fe ON u.id = fe.following_id AND fe.status = 'accept'
		LEFT JOIN
			post p ON p.author_id = u.id
		LEFT JOIN
			post_like pl ON pl.id = p.id

		WHERE
			u.id != %d AND
			NOT EXISTS (
				SELECT 1
				FROM follower
				WHERE (follower_id = u.id AND following_id = %d AND status = 'accept')
					OR (follower_id = %d AND following_id = u.id AND status = 'accept')
			)
		GROUP BY
			%s, u.created_at
		LIMIT '%s' OFFSET '%s'
	`

	// on joining websocket
	QUERY_ON_WS_PAYLOAD = `
	SELECT (
		SELECT json_group_array(
			json_object(
				'id', u.id,
				'first_name', u.first_name,
				'last_name', u.last_name,
				'email', u.email,
				'username', u.username,
				'avatar', u.avatar,
				'unread_message', (
					SELECT COUNT(*)
					FROM privateMessage pm
					WHERE (pm.sender_id = u.id AND pm.receiver_id = ?)
					AND pm.status = 'unread'
				),
				'last_activity', (
					SELECT MAX(m.created_at)
					FROM privateMessage m
					WHERE (m.sender_id = u.id AND m.receiver_id = ?) OR
					(m.sender_id = ? OR m.receiver_id = u.id)
				),
				'online', CASE WHEN u.id IN (%s) THEN 1 ELSE 0 END
			)
		)
		FROM (
			SELECT DISTINCT u.id, u.first_name, u.last_name, u.email, u.username, u.avatar
			FROM user u
			LEFT JOIN follower f ON (f.follower_id = ? OR f.following_id = ?)
			WHERE (u.id = f.follower_id AND f.status = 'accept' OR u.id = f.following_id AND f.status = 'accept') AND u.id != ?
		) u
		ORDER BY 'last_activity' DESC
	) AS users,
		(
			SELECT json_group_array(
				json_object(
					'id', g.id,
					'title', g.title,
					'description', g.description,

					'author_id', g.author_id,
					'author_username', u.username,
                    'author_first_name', u.first_name,
                    'author_last_name', u.last_name,

                    'author_avatar', u.avatar,
					'cover', g.cover,
					'created_at', g.created_at,

					'member_count', (
                        SELECT COUNT(*)
                        FROM groupMembers gm
                        WHERE gm.group_id = g.id AND gm.status = 'accepted'
                    ),

					'post_count', (
                        SELECT COUNT(*)
                        FROM groupPosts gp
                        WHERE gp.group_id = g.id
                    ),

					'event_count', (
                        SELECT COUNT(*)
                        FROM groupEvents ge
                        WHERE ge.group_id = g.id
                    ),

					'unread_message', (
						SELECT COUNT(*)
						FROM groupMessageStatus gmsgs
						WHERE
						gmsgs.group_message_id IN (SELECT gmsg.id FROM groupMessage gmsg WHERE gmsg.group_id = g.id)
						AND gmsgs.user_id = ?
						AND gmsgs.status = 'unread'
					),
					'last_activity', (
						SELECT MAX(last_time)
						FROM (
							SELECT MAX(gmsg.created_at) AS last_time FROM groupMessage gmsg WHERE gmsg.group_id = g.id
							UNION
							SELECT MAX(ge.created_at) AS last_time FROM groupEvents ge WHERE ge.group_id = g.id
							UNION
							SELECT MAX(gp.created_at) AS last_time FROM groupPosts gp WHERE gp.group_id = g.id
						) AS last_time
					),
					'is_part_of_group', 1
				)
			)
			FROM groups g
			LEFT JOIN groupMembers gm ON gm.user_id = ? AND gm.status = 'accepted'
			LEFT JOIN user u ON u.id = g.author_id
			WHERE gm.group_id = g.id
			ORDER BY 'last_activity' DESC
		) AS groups,
		(
			SELECT COUNT(*)
			FROM notifications
			WHERE receiver_id = ?
			AND status = 'unread'
		) AS unread_notifications;
	`

	QUERY_ON_WS_PAYLOAD_USERS = `
		SELECT
		(
			SELECT json_group_array(
				json_object(
					'id', u.id,
					'first_name', u.first_name,
					'last_name', u.last_name,
					'email', u.email,
					'username', u.username,
					'avatar', u.avatar,
					'unread_message', (
						SELECT COUNT(*)
						FROM privateMessage pm
						WHERE (pm.sender_id = u.id AND pm.receiver_id = %d)
						AND pm.status = 'unread'
					),
					'last_activity', (
						SELECT MAX(m.created_at)
						FROM privateMessage m
						WHERE (m.sender_id = u.id AND m.receiver_id = %d) OR
							(m.sender_id = %d OR m.receiver_id = u.id)
					),
					"online", CASE WHEN u.id IN (%s) THEN 1 ELSE 0 END
				)
			)
			FROM (
				SELECT DISTINCT u.id, u.first_name, u.last_name, u.email, u.username, u.avatar
				FROM user u
				LEFT JOIN follower f ON (f.follower_id = %d OR f.following_id = %d)
				WHERE (u.id = f.follower_id AND f.status = 'accept' OR u.id = f.following_id AND f.status = 'accept') AND u.id != %d
			) u
			ORDER BY 'last_activity' DESC
		) AS users
	`

	QUERY_ON_WS_PAYLOAD_GROUPS = `
		SELECT
		(
			SELECT json_group_array(
				json_object(
					'id', g.id,
					'title', g.title,
					'description', g.description,

					'author_id', g.author_id,
					'author_username', u.username,
                    'author_first_name', u.first_name,
                    'author_last_name', u.last_name,

                    'author_avatar', u.avatar,
					'cover', g.cover,
					'created_at', g.created_at, 

					'member_count', (
                        SELECT COUNT(*)
                        FROM groupMembers gm
                        WHERE gm.group_id = g.id AND gm.status = 'accepted'
                    ),

					'post_count', (
                        SELECT COUNT(*)
                        FROM groupPosts gp
                        WHERE gp.group_id = g.id
                    ),

					'event_count', (
                        SELECT COUNT(*)
                        FROM groupEvents ge
                        WHERE ge.group_id = g.id
                    ),

					'unread_message', (
						SELECT COUNT(*)
						FROM groupMessageStatus gmsgs
						WHERE
						gmsgs.group_message_id IN (SELECT gmsg.id FROM groupMessage gmsg WHERE gmsg.group_id = g.id)
						AND gmsgs.user_id = %d
						AND gmsgs.status = 'unread'
					),
					'last_activity', (
						SELECT MAX(last_time)
						FROM (
							SELECT MAX(gmsg.created_at) AS last_time FROM groupMessage gmsg WHERE gmsg.group_id = g.id
							UNION
							SELECT MAX(ge.created_at) AS last_time FROM groupEvents ge WHERE ge.group_id = g.id
							UNION
							SELECT MAX(gp.created_at) AS last_time FROM groupPosts gp WHERE gp.group_id = g.id
						) AS last_time
					),
					'is_part_of_group', 1
				)
			)
			FROM groups g
			LEFT JOIN groupMembers gm ON gm.user_id = %d AND gm.status = 'accepted'
			LEFT JOIN user u ON u.id = g.author_id
			WHERE gm.group_id = g.id
			ORDER BY 'last_activity' DESC
		) AS groups;
	`

	QUERY_ON_WS_PAYLOAD_NOTIFICATIONS = `
		SELECT
		(
			SELECT COUNT(*)
			FROM notifications
			WHERE receiver_id = %d
			AND status = 'unread'
		) AS unread_notifications;

	`

	// get all follower
	QUERY_GETTING_ALL_FOLLOWERS = `
		SELECT
			u.id,
			u.username,
			u.email,
			u.first_name,
			u.last_name,
			u.avatar,
			f.created_at

		FROM
			follower f
		LEFT JOIN
			user u ON u.id = f.following_id

		WHERE
			f.follower_id = %d AND f.status = 'accept'
		ORDER BY
			f.created_at DESC;
	`

	// get all users following
	QUERY_GETTING_ALL_FOLLOWING = `
		SELECT
			u.id,
			u.username,
			u.email,
			u.first_name,
			u.last_name,
			u.avatar,
			f.created_at

		FROM
			follower f
		LEFT JOIN
			user u ON u.id = f.follower_id

		WHERE
			f.following_id = %d AND f.status = 'accept'
		ORDER BY
			f.created_at DESC;
	`

	// Private messages
	QUERY_GETTING_PRIVATE_MESSAGE = `SELECT
		pm.id,
		pm.sender_id,
		pm.receiver_id,
		pm.content,

		sender.first_name AS sender_first_name,
		sender.last_name AS sender_last_name,
		sender.username AS sender_username,
		sender.avatar AS sender_avatar,

		receiver.first_name AS receiver_first_name,
		receiver.last_name AS receiver_last_name,
		receiver.username AS receiver_username,
		receiver.avatar AS receiver_avatar,

		pm.created_at

		FROM privateMessage pm
		LEFT JOIN user sender ON pm.sender_id = sender.id
		LEFT JOIN user receiver ON pm.receiver_id = receiver.id

		WHERE pm.id = %d
	`

	// all Private messages for a user
	QUERYGETTINGALLPRIVATEMESSAGE = `SELECT
		pm.id,
		pm.sender_id,
		pm.receiver_id,
		pm.content,

		sender.first_name AS sender_first_name,
		sender.last_name AS sender_last_name,
		sender.username AS sender_username,
		sender.avatar AS sender_avatar,

		receiver.first_name AS receiver_first_name,
		receiver.last_name AS receiver_last_name,
		receiver.username AS receiver_username,
		receiver.avatar AS receiver_avatar,

		pm.created_at

		FROM
			privateMessage pm
		LEFT JOIN
			user sender ON pm.sender_id = sender.id
		LEFT JOIN
			user receiver ON pm.receiver_id = receiver.id

		WHERE
			(pm.sender_id = %d AND pm.receiver_id = %d) OR (pm.sender_id = %d AND pm.receiver_id = %d)
		ORDER BY
			pm.created_at DESC
		LIMIT %s OFFSET %s
	`

	// Groups messages
	QUERY_GETTING_GROUPS_MESSAGE = `SELECT
		gm.id,
		gm.sender_id,
		gm.group_id,
		gm.content,
		gm.created_at,

		sender.first_name AS sender_first_name,
		sender.last_name AS sender_last_name,
		sender.username AS sender_username,
		sender.avatar AS sender_avatar,

		g.title AS group_title,
		g.description AS group_description,
		g.cover AS group_cover,
		g.created_at AS group_created_at

		FROM
			groupMessage gm
		LEFT
			JOIN user sender ON gm.sender_id = sender.id
		LEFT JOIN
			groups g ON gm.group_id = g.id

		WHERE gm.id = %d
	`

	// get all groups messages
	QUERY_GETTING_ALL_GROUPS_MESSAGE = `SELECT
		gm.id,
		gm.sender_id,
		gm.group_id,
		gm.content,
		gm.created_at,

		sender.first_name AS sender_first_name,
		sender.last_name AS sender_last_name,
		sender.username AS sender_username,
		sender.avatar AS sender_avatar,

		g.title AS group_title,
		g.description AS group_description,
		g.cover AS group_cover,
		g.created_at AS group_created_at

		FROM
			groupMessage gm
		LEFT JOIN
			user sender ON gm.sender_id = sender.id
		LEFT JOIN
			groups g ON gm.group_id = g.id

		WHERE
			gm.group_id = %d
		ORDER BY
			gm.created_at DESC
		LIMIT %s OFFSET %s
	`

	// FOLLOW notifications
	// all
	QUERY_GETTING_ALL_FOLLOW_REQUEST_NOTIFICATION = `SELECT
		n.id,
		n.sender_id,
		n.receiver_id,
		n.notification_type,

		sender.first_name AS sender_first_name,
		sender.last_name AS sender_last_name,
		sender.username AS sender_username,
		sender.avatar AS sender_avatar,

		receiver.first_name AS receiver_first_name,
		receiver.last_name AS receiver_last_name,
		receiver.username AS receiver_username,
		receiver.avatar AS receiver_avatar

		FROM notifications n
		LEFT JOIN user sender ON n.sender_id = sender.id
		LEFT JOIN user receiver ON n.receiver_id = receiver.id

		WHERE n.notification_type = 'follow_request' AND n.receiver_id = '%d'
		LIMIT '%s' OFFSET '%s'
	`

	// single
	QUERY_GETTING_FOLLOW_REQUEST_NOTIFICATION = `SELECT
		n.id,
		n.sender_id,
		n.receiver_id,
		n.notification_type,

		sender.first_name AS sender_first_name,
		sender.last_name AS sender_last_name,
		sender.username AS sender_username,
		sender.avatar AS sender_avatar,

		receiver.first_name AS receiver_first_name,
		receiver.last_name AS receiver_last_name,
		receiver.username AS receiver_username,
		receiver.avatar AS receiver_avatar

		FROM notifications n
		LEFT JOIN user sender ON n.sender_id = sender.id
		LEFT JOIN user receiver ON n.receiver_id = receiver.id

		WHERE n.notification_type = 'follow_request' AND n.id = '%d' AND n.receiver_id = '%d'
	`

	// INVITATION and REQUEST group notifications
	// all
	QUERY_GETTING_ALL_GROUPS_INVITED_REQUESTED_NOTIFICATION = `SELECT
		n.id,
		n.sender_id,
		n.receiver_id,
		n.notification_type,

		sender.first_name AS sender_first_name,
		sender.last_name AS sender_last_name,
		sender.username AS sender_username,
		sender.avatar AS sender_avatar,

		receiver.first_name AS receiver_first_name,
		receiver.last_name AS receiver_last_name,
		receiver.username AS receiver_username,
		receiver.avatar AS receiver_avatar,

		g.id AS group_id,
		g.title AS title,
		g.description AS description

		FROM notifications n
		LEFT JOIN user sender ON n.sender_id = sender.id
		LEFT JOIN user receiver ON n.receiver_id = receiver.id
		LEFT JOIN groups g ON g.id = n.group_id

		WHERE n.notification_type IN ('groups_invited', 'groups_requested') AND n.receiver_id = '%d'
		LIMIT '%s' OFFSET '%s'
	`

	// single
	QUERY_GETTING_GROUPS_INVITED_REQUESTED_NOTIFICATION = `SELECT
		n.id,
		n.sender_id,
		n.receiver_id,
		n.notification_type,

		sender.first_name AS sender_first_name,
		sender.last_name AS sender_last_name,
		sender.username AS sender_username,
		sender.avatar AS sender_avatar,

		receiver.first_name AS receiver_first_name,
		receiver.last_name AS receiver_last_name,
		receiver.username AS receiver_username,
		receiver.avatar AS receiver_avatar,

		g.id AS group_id,
		g.title AS title,
		g.description AS description

		FROM notifications n
		LEFT JOIN user sender ON n.sender_id = sender.id
		LEFT JOIN user receiver ON n.receiver_id = receiver.id
		LEFT JOIN groups g ON g.id = n.group_id

		WHERE n.id = '%d' AND n.notification_type IN ('groups_invited', 'groups_requested') AND n.receiver_id = '%d'
	`

	// EVENT notification
	// all
	QUERY_GETTING_ALL_GROUP_EVENTS_NOTIFICATION = `SELECT
		n.id,
		n.sender_id,
		n.receiver_id,
		n.notification_type,

		sender.first_name AS sender_first_name,
		sender.last_name AS sender_last_name,
		sender.username AS sender_username,
		sender.avatar AS sender_avatar,

		receiver.first_name AS receiver_first_name,
		receiver.last_name AS receiver_last_name,
		receiver.username AS receiver_username,
		receiver.avatar AS receiver_avatar,

		n.group_id AS group_id,
		g.title AS group_title,
		g.description AS group_description,

		n.event_id AS event_id,
		e.title AS event_title,
		e.description AS event_description,
		e.author_id AS author_id,
		e.datetime,

		author.first_name AS author_first_name,
		author.last_name AS author_last_name,
		author.username AS author_username,
		author.avatar AS author_avatar

		FROM notifications n
		LEFT JOIN user sender ON n.sender_id = sender.id
		LEFT JOIN user receiver ON n.receiver_id = receiver.id
		LEFT JOIN groups g ON g.id = n.group_id
		LEFT JOIN groupEvents e ON e.id = n.event_id
		LEFT JOIN user author ON e.author_id = author.id

		WHERE n.notification_type = 'groups_events' AND n.receiver_id = '%d'
		LIMIT '%s' OFFSET '%s'
	`

	// single
	QUERY_GETTING_GROUP_EVENTS_NOTIFICATION = `SELECT
		n.id,
		n.sender_id,
		n.receiver_id,
		n.notification_type,

		sender.first_name AS sender_first_name,
		sender.last_name AS sender_last_name,
		sender.username AS sender_username,
		sender.avatar AS sender_avatar,

		receiver.first_name AS receiver_first_name,
		receiver.last_name AS receiver_last_name,
		receiver.username AS receiver_username,
		receiver.avatar AS receiver_avatar,

		g.id AS group_id,
		g.title AS group_title,
		g.description AS group_description,

		e.id,
		e.title AS event_title,
		e.description AS event_description,
		e.author_id AS author_id,
		e.datetime,

		author.first_name AS author_first_name,
		author.last_name AS author_last_name,
		author.username AS author_username,
		author.avatar AS author_avatar

		FROM notifications n
		LEFT JOIN user sender ON n.sender_id = sender.id
		LEFT JOIN user receiver ON n.receiver_id = receiver.id
		LEFT JOIN groups g ON g.id = n.group_id
		LEFT JOIN groupEvents e ON e.group_id = g.id
		LEFT JOIN user author ON e.author_id = author.id

		WHERE n.notification_type = 'groups_events' AND n.id = %d AND n.receiver_id = '%d'
	`
)
