<?php
final class WP_Site
{
	public $blog_id;
	public $domain = '';
	public $path = '';
	public $site_id = '0';
	public $registered = '0000-00-00 00:00:00';
	public $last_updated = '0000-00-00 00:00:00';
	public $public = '1';
	public $archived = '0';
	public $mature = '0';
	public $spam = '0';
	public $deleted = '0';
	public $lang_id = '0';
	public static function get_instance($site_id)
	{
		global $wpdb;
		$site_id = (int)$site_id;
		if (!$site_id) {
			return false;
		}
		$_site = wp_cache_get($site_id, 'sites');
		if (!$_site) {
			$_site = $wpdb->get_row($wpdb->prepare("SELECT * FROM {$wpdb->blogs} WHERE blog_id = %d LIMIT 1", $site_id));
			if (empty($_site) || is_wp_error($_site)) {
				return false;
			}
			wp_cache_add($site_id, $_site, 'sites');
		}
		return new WP_Site($_site);
	}
	public function __construct($site)
	{
		foreach (get_object_vars($site) as $key => $value) {
			$this->$key = $value;
		}
	}
	public function to_array()
	{
		return get_object_vars($this);
	}
	public function __get($key)
	{
		switch ($key) {
			case 'id':
				return (int)$this->blog_id;
			case 'network_id':
				return (int)$this->site_id;
			case 'blogname':
			case 'siteurl':
			case 'post_count':
			case 'home':
			default:
				if (!did_action('ms_loaded')) {
					return null;
				}
				$details = $this->get_details();
				if (isset($details->$key)) {
					return $details->$key;
				}
		}
		return null;
	}
	public function __isset($key)
	{
		switch ($key) {
			case 'id':
			case 'network_id':
				return true;
			case 'blogname':
			case 'siteurl':
			case 'post_count':
			case 'home':
				if (!did_action('ms_loaded')) {
					return false;
				}
				return true;
			default:
				if (!did_action('ms_loaded')) {
					return false;
				}
				$details = $this->get_details();
				if (isset($details->$key)) {
					return true;
				}
		}
		return false;
	}
	public function __set($key, $value)
	{
		switch ($key) {
			case 'id':
				$this->blog_id = (string)$value;
				break;
			case 'network_id':
				$this->site_id = (string)$value;
				break;
			default:
				$this->$key = $value;
		}
	}
	private function get_details()
	{
		$details = wp_cache_get($this->blog_id, 'site-details');
		if (false === $details) {
			switch_to_blog($this->blog_id);
			$details = new stdClass();
			foreach (get_object_vars($this) as $key => $value) {
				$details->$key = $value;
			}
			$details->blogname = get_option('blogname');
			$details->siteurl = get_option('siteurl');
			$details->post_count = get_option('post_count');
			$details->home = get_option('home');
			restore_current_blog();
			wp_cache_set($this->blog_id, $details, 'site-details');
		}
		$details = apply_filters_deprecated('blog_details', [$details], '4.7.0', 'site_details');
		$details = apply_filters('site_details', $details);
		return $details;
	}
}

